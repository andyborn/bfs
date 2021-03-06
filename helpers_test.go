package bfs_test

import (
	"context"

	"github.com/bsm/bfs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helpers", func() {
	var bucket *bfs.InMem
	var ctx = context.Background()

	BeforeEach(func() {
		bucket = bfs.NewInMem()
	})

	It("should write objects", func() {
		err := bfs.WriteObject(bucket, ctx, "path/to/file", []byte("testdata"))
		Expect(err).NotTo(HaveOccurred())

		Expect(bucket.ObjectSizes()).
			To(HaveKeyWithValue("path/to/file", int64(8)))
	})

	It("should copy objects", func() {
		err := bfs.WriteObject(bucket, ctx, "src.txt", []byte("testdata"))
		Expect(err).NotTo(HaveOccurred())

		err = bfs.CopyObject(bucket, ctx, "src.txt", "dst.txt")
		Expect(err).NotTo(HaveOccurred())

		Expect(bucket.ObjectSizes()).
			To(HaveKeyWithValue("src.txt", int64(8)))
		Expect(bucket.ObjectSizes()).
			To(HaveKeyWithValue("dst.txt", int64(8)))
	})
})
