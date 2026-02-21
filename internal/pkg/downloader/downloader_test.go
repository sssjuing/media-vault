package downloader

import "testing"

func createDownloader(t *testing.T) *Downloader {
	url := "https://demo.unified-streaming.com/k8s/features/stable/video/tears-of-steel/tears-of-steel.ism/tears-of-steel-audio_eng=64008-video_eng=401000.m3u8"
	name := "tears-of-steel"

	r := NewResource(url, name, "tmp", nil)
	sf := WithSegmentFinish(func(sr *SegmentRow, index int) {
		t.Logf("download %d segment %s finished, status: %d", index, sr.Path, sr.Status)
	})
	tf := WithFinally(func(r *Resource, err error) {
		if err != nil {
			t.Log("task error:", err)
		}
		t.Logf("task %s finished", r.Filename)
	})
	return New(r, sf, tf)
}

func TestMakeSegmentList(t *testing.T) {
	d := createDownloader(t)
	d.makeSegmentList()
	t.Logf("name: %s; segment length: %d", d.resource.Filename, len(d.resource.SegmentList))
	t.Log(d.resource.SegmentList[0])
}

func TestDownloadSegments(t *testing.T) {
	d := createDownloader(t)
	d.makeSegmentList()
	d.resource.SegmentList = d.resource.SegmentList[:5]
	d.resource.SegmentList[2].Status = 1

	d.makeTempDir()
	d.downloadSegments(func(sr *SegmentRow, index int) {
		t.Log(index, sr.Path)
	})
	for idx, sr := range d.resource.SegmentList {
		t.Log(idx, sr.Path, sr.Status)
	}
}

func TestMergeSegments(t *testing.T) {
	d := createDownloader(t)
	d.makeSegmentList()
	d.resource.SegmentList = d.resource.SegmentList[:5]

	d.makeTempDir()
	d.downloadSegments(func(sr *SegmentRow, index int) {
		t.Log(sr.Path)
	})
	d.mergeSegments("tmp")
}
