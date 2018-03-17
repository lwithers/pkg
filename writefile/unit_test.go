package writefile

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	tmpdir       string
	tmpfileCount int
)

func TestMain(m *testing.M) {
	var err error
	if tmpdir, err = ioutil.TempDir("", "tmpfile-unit-test"); err != nil {
		fmt.Fprintf(os.Stderr, "Could not create temporary "+
			"directory: %v\n", err)
		os.Exit(1)
	}

	ret := m.Run()

	os.RemoveAll(tmpdir)
	os.Exit(ret)
}

func TestCommit(t *testing.T) {
	expData := []byte("test")

	ff, tmpf := makeTempFile(t, "commit", "commit")
	if _, err := tmpf.Write(expData); err != nil {
		t.Errorf("error writing to temporary file: %v", err)
	}
	if err := Commit(ff, tmpf); err != nil {
		t.Errorf("error committing temporary file: %v", err)
	}

	checkFile(t, "final", ff, expData)
}

func TestOverwrite(t *testing.T) {
	origData := []byte("original")
	expData := []byte("testing")

	// write out original data
	ff := filepath.Join(tmpdir, "overwrite")
	if err := ioutil.WriteFile(ff, origData, 0666); err != nil {
		t.Fatalf("could not prepare file %s: %v", ff, err)
	}

	// open a temporary file — should not overwrite our original
	_, tmpf := makeTempFile(t, "overwrite", "overwrite")
	checkFile(t, "tmpfile open", ff, origData)

	// write some data to temporary file, and verify original still intact
	if _, err := tmpf.Write(expData); err != nil {
		t.Errorf("error writing to temporary file: %v", err)
	}
	checkFile(t, "tmpfile modified", ff, origData)

	// now Commit — original should be overwritten
	if err := Commit(ff, tmpf); err != nil {
		t.Errorf("error committing temporary file: %v", err)
	}
	checkFile(t, "tmpfile committed", ff, expData)
}

func TestAbort(t *testing.T) {
	origData := []byte("original")
	expData := []byte("testing")

	// write out original data
	ff := filepath.Join(tmpdir, "overwrite_abort")
	if err := ioutil.WriteFile(ff, origData, 0666); err != nil {
		t.Fatalf("could not prepare file %s: %v", ff, err)
	}

	// open a temporary file — should not overwrite our original
	_, tmpf := makeTempFile(t, "overwrite_abort", "overwrite_abort")
	checkFile(t, "tmpfile open", ff, origData)

	// write some data to temporary file, and verify original still intact
	if _, err := tmpf.Write(expData); err != nil {
		t.Errorf("error writing to temporary file: %v", err)
	}
	checkFile(t, "tmpfile modified", ff, origData)

	// now Abort — original should not be overwritten
	Abort(tmpf)
	checkFile(t, "tmpfile aborted", ff, origData)
	if err := tmpf.Close(); err == nil {
		t.Error("expected error when closing aborted file")
	}
}

func TestSymlinkRel(t *testing.T) {
	origData := []byte("original")
	expData := []byte("testing")

	// write out original data
	ff := filepath.Join(tmpdir, "sym_target")
	if err := ioutil.WriteFile(ff, origData, 0666); err != nil {
		t.Fatalf("could not prepare file %s: %v", ff, err)
	}

	// create a symlink
	tgt := filepath.Join(tmpdir, "sym")
	if err := os.Symlink("sym_target", tgt); err != nil {
		t.Fatalf("could not prepare symlink %s: %v", tgt, err)
	}

	// open a temporary file — should not overwrite our original
	_, tmpf := makeTempFile(t, "sym", "sym_target")
	checkFile(t, "tmpfile open", ff, origData)

	// write some data to temporary file, and verify original still intact
	if _, err := tmpf.Write(expData); err != nil {
		t.Errorf("error writing to temporary file: %v", err)
	}
	checkFile(t, "tmpfile modified", ff, origData)

	// now Commit — original should be overwritten
	if err := Commit(ff, tmpf); err != nil {
		t.Errorf("error committing temporary file: %v", err)
	}
	checkFile(t, "tmpfile committed", ff, expData)
}

func TestSymlinkAbs(t *testing.T) {
	origData := []byte("original")
	expData := []byte("testing")

	// write out original data
	ff := filepath.Join(tmpdir, "sym_target_abs")
	if err := ioutil.WriteFile(ff, origData, 0666); err != nil {
		t.Fatalf("could not prepare file %s: %v", ff, err)
	}

	// create a symlink
	tgt := filepath.Join(tmpdir, "sym_abs")
	if err := os.Symlink(filepath.Join(tmpdir, "sym_target_abs"), tgt); err != nil {
		t.Fatalf("could not prepare symlink %s: %v", tgt, err)
	}

	// open a temporary file — should not overwrite our original
	_, tmpf := makeTempFile(t, "sym_abs", "sym_target_abs")
	checkFile(t, "tmpfile open", ff, origData)

	// write some data to temporary file, and verify original still intact
	if _, err := tmpf.Write(expData); err != nil {
		t.Errorf("error writing to temporary file: %v", err)
	}
	checkFile(t, "tmpfile modified", ff, origData)

	// now Commit — original should be overwritten
	if err := Commit(ff, tmpf); err != nil {
		t.Errorf("error committing temporary file: %v", err)
	}
	checkFile(t, "tmpfile committed", ff, expData)
}

func TestSymlinkEnoent(t *testing.T) {
	expData := []byte("testing")

	// create a symlink (pointing at a file which doesn't exist)
	tgt := filepath.Join(tmpdir, "sym_enoent")
	if err := os.Symlink("enoent", tgt); err != nil {
		t.Fatalf("could not prepare symlink %s: %v", tgt, err)
	}

	// open a temporary file
	ff, tmpf := makeTempFile(t, "sym_enoent", "enoent")

	// write some data to temporary file, commit and verify
	if _, err := tmpf.Write(expData); err != nil {
		t.Errorf("error writing to temporary file: %v", err)
	}
	if err := Commit(ff, tmpf); err != nil {
		t.Errorf("error committing temporary file: %v", err)
	}
	checkFile(t, "tmpfile committed", ff, expData)
}

func TestSymlinkLoop(t *testing.T) {
	a := filepath.Join(tmpdir, "symlink_loop_a")
	b := filepath.Join(tmpdir, "symlink_loop_b")
	if err := os.Symlink(a, b); err != nil {
		t.Fatalf("failed to create symlink %s→%s: %v", a, b, err)
	}
	if err := os.Symlink(b, a); err != nil {
		t.Fatalf("failed to create symlink %s→%s: %v", b, a, err)
	}
	ff, tmpf, err := New(a)
	switch err := err.(type) {
	case nil:
		t.Errorf("error expected, but New succeeded (%q/%q)",
			ff, tmpf.Name())
	case *os.PathError:
		if err.Path != a {
			t.Errorf("got *os.PathError with unexpected path "+
				"(got %#v, expected %q)", err, a)
		}
	default:
		t.Errorf("got unexpected error type %T (%v)", err, err)
	}
}

func TestSymlinkDir(t *testing.T) {
	dnam := filepath.Join(tmpdir, "target_dir")
	if err := os.Mkdir(dnam, 0777); err != nil {
		t.Fatal(err)
	}

	sym := filepath.Join(tmpdir, "sym_dir")
	if err := os.Symlink(dnam, sym); err != nil {
		t.Fatal(err)
	}

	ff, tmpf, err := New(sym)
	switch err := err.(type) {
	case nil:
		t.Errorf("error expected, but New succeeded (%q/%q)",
			ff, tmpf.Name())
	case *os.PathError:
		if err.Path != sym {
			t.Errorf("got *os.PathError with unexpected path "+
				"(got %#v, expected %q)", err, sym)
		}
	default:
		t.Errorf("got unexpected error type %T (%v)", err, err)
	}
}

func makeTempFile(t *testing.T, srcName, expFinal string) (string, *os.File) {
	tmpfileCount++
	final := filepath.Join(tmpdir, srcName)
	ff, tmpf, err := New(final)
	if err != nil {
		t.Fatalf("could not create temporary file for %q: %v",
			final, err)
	}
	if ff != filepath.Join(tmpdir, expFinal) {
		t.Fatalf("unexpected final fname %q does not match expected %q",
			ff, filepath.Join(tmpdir, expFinal))
	}
	if filepath.Dir(tmpf.Name()) != tmpdir {
		t.Fatalf("unexpected final dir %q does not match expected %q",
			filepath.Dir(tmpf.Name()), tmpdir)
	}
	return ff, tmpf
}

func checkFile(t *testing.T, label, fname string, exp []byte) {
	if raw, err := ioutil.ReadFile(fname); err != nil {
		t.Errorf("%s: could not read file: %v", label, err)
	} else if !bytes.Equal(raw, exp) {
		t.Errorf("%s: data in %s not as expected", label, fname)
		t.Errorf("%s: got %q, expected %q", label, raw, exp)
	}
}
