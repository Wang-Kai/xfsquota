package mountpoint

import "testing"

func TestParseMountInfoLine(t *testing.T) {
	mountInfoLine := `76 86 254:16 /a /data/a rw,relatime shared:28 - xfs /dev/vdb rw,attr2,inode64,prjquota`
	mnt := parseMountInfoLine(mountInfoLine)
	t.Logf("%+v", mnt)
}
