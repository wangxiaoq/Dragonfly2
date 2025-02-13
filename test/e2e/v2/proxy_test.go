/*
 *     Copyright 2024 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package e2e

import (
	"fmt"
	"time"

	. "github.com/onsi/ginkgo/v2" //nolint
	. "github.com/onsi/gomega"    //nolint

	"d7y.io/dragonfly/v2/test/e2e/v2/util"
)

var _ = Describe("Download Using Proxy", func() {
	Context("/etc/containerd/config.toml file", func() {
		It("download should be ok", Label("proxy", "download"), func() {
			clientPod, err := util.ClientExec()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := clientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -H 'X-Dragonfly-Tag: proxy' %s --output %s", util.GetFileURL("/etc/containerd/config.toml"), util.GetOutputPath("config-proxy.coml"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "0ef882badaeb3195aed759a203cfb61951d158bf614f90ab0a20504fb7f97992",
				Sha256: "66404431f9a0d5c78205e5a3eb041f76767094fc278c0a091d4ffa10f06cf641",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{clientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{clientPod}, util.GetOutputPath("config-proxy.coml"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(1 * time.Second)
			seedClientPods := make([]*util.PodExec, 3)
			for i := 0; i < 3; i++ {
				seedClientPods[i], err = util.SeedClientExec(i)
				fmt.Println(err)
				Expect(err).NotTo(HaveOccurred())
			}

			sha256sum, err = util.CalculateSha256ByTaskID(seedClientPods, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file", func() {
		It("download should be ok", Label("proxy", "download"), func() {
			clientPod, err := util.ClientExec()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := clientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -H 'X-Dragonfly-Tag: proxy' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-proxy"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "e37fccc0c725a947c0a8856e2e1f6a14a4a5792338a73dcafa7e5ebd6443f7b4",
				Sha256: "327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{clientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{clientPod}, util.GetOutputPath("kubectl-proxy"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(1 * time.Second)
			seedClientPods := make([]*util.PodExec, 3)
			for i := 0; i < 3; i++ {
				seedClientPods[i], err = util.SeedClientExec(i)
				fmt.Println(err)
				Expect(err).NotTo(HaveOccurred())
			}

			sha256sum, err = util.CalculateSha256ByTaskID(seedClientPods, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))
		})
	})

	Context("/bin/x86_64 file", func() {
		It("download should be ok", Label("proxy", "download"), func() {
			clientPod, err := util.ClientExec()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := clientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -H 'X-Dragonfly-Tag: proxy' %s --output %s", util.GetFileURL("/bin/x86_64"), util.GetOutputPath("x86_64-proxy"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "7c7cf5c8303626754ba4ae6575c1c96cd3fcad985687ec0d15744e68b15661d6",
				Sha256: "a1cbf1bf2d66757121677fd7fefafacd4f843a2cb44a451131002803bae56a65",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{clientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{clientPod}, util.GetOutputPath("x86_64-proxy"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(1 * time.Second)
			seedClientPods := make([]*util.PodExec, 3)
			for i := 0; i < 3; i++ {
				seedClientPods[i], err = util.SeedClientExec(i)
				fmt.Println(err)
				Expect(err).NotTo(HaveOccurred())
			}

			sha256sum, err = util.CalculateSha256ByTaskID(seedClientPods, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))
		})
	})

	Context("/bin/zless file", func() {
		It("download should be ok", Label("proxy", "download"), func() {
			clientPod, err := util.ClientExec()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := clientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -H 'X-Dragonfly-Tag: proxy' %s --output %s", util.GetFileURL("/bin/zless"), util.GetOutputPath("zless-proxy"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "297205e16a85e8cf803067c84fdd659c1877ff595e76e7633e4aa52b64321dee",
				Sha256: "b0cfe211f851049a78f5812cf5b1d7cb4f3fbb101c02c1865c940d5797f4b73b",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{clientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{clientPod}, util.GetOutputPath("zless-proxy"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(1 * time.Second)
			seedClientPods := make([]*util.PodExec, 3)
			for i := 0; i < 3; i++ {
				seedClientPods[i], err = util.SeedClientExec(i)
				fmt.Println(err)
				Expect(err).NotTo(HaveOccurred())
			}

			sha256sum, err = util.CalculateSha256ByTaskID(seedClientPods, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))
		})
	})

	Context("/bin/bash file", func() {
		It("download should be ok", Label("proxy", "download"), func() {
			clientPod, err := util.ClientExec()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := clientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -H 'X-Dragonfly-Tag: proxy' %s --output %s", util.GetFileURL("/bin/bash"), util.GetOutputPath("bash-proxy"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "b4df90b001f49adcc64d71f68b1c01ad462f631c94c7c7ffc12d97f351874da1",
				Sha256: "c37f93c73cf2f303f874c094f6f76e47b2421a3da9f0e7e0b98bea8a3d685322",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{clientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{clientPod}, util.GetOutputPath("bash-proxy"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(1 * time.Second)
			seedClientPods := make([]*util.PodExec, 3)
			for i := 0; i < 3; i++ {
				seedClientPods[i], err = util.SeedClientExec(i)
				fmt.Println(err)
				Expect(err).NotTo(HaveOccurred())
			}

			sha256sum, err = util.CalculateSha256ByTaskID(seedClientPods, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("proxy", "download", "range: bytes=-100"), func() {
			clientPod, err := util.ClientExec()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := clientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r -100 -H 'X-Dragonfly-Tag: proxy-bytes-100' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-proxy-bytes-100"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "b56a0807131865baffe2d687440f121f40cf2071c478a5abfb0930bc57b7e715",
				Sha256: "cd00e292c5970d3c5e2f0ffa5171e555bc46bfc4faddfb4a418b6840b86e79a3",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{clientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("81d9f55e659a1b764e99dec6915ed47aa66955d30c47ffd4d427b4194a31684d").To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{clientPod}, util.GetOutputPath("kubectl-proxy-bytes-100"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(1 * time.Second)
			seedClientPods := make([]*util.PodExec, 3)
			for i := 0; i < 3; i++ {
				seedClientPods[i], err = util.SeedClientExec(i)
				fmt.Println(err)
				Expect(err).NotTo(HaveOccurred())
			}

			sha256sum, err = util.CalculateSha256ByTaskID(seedClientPods, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("81d9f55e659a1b764e99dec6915ed47aa66955d30c47ffd4d427b4194a31684d").To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("proxy", "download", "range: bytes=0-100"), func() {
			clientPod, err := util.ClientExec()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := clientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r 0-100 -H 'X-Dragonfly-Tag: proxy-bytes-0-100' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-proxy-bytes-0-100"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "0091f6a0373fe02d2e4cf6a83904889f5326a3590239015743bc812b02c51408",
				Sha256: "7bbfcb694f6acc69483751d3b48d5bdbdcb284f9e04c7b7caa04b5977cffddc8",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{clientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("0dd6bcd036002bbd79cd8dfebcd9813ec4231562c0b15b02cbc7c70a062717f9").To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{clientPod}, util.GetOutputPath("kubectl-proxy-bytes-0-100"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(1 * time.Second)
			seedClientPods := make([]*util.PodExec, 3)
			for i := 0; i < 3; i++ {
				seedClientPods[i], err = util.SeedClientExec(i)
				fmt.Println(err)
				Expect(err).NotTo(HaveOccurred())
			}

			sha256sum, err = util.CalculateSha256ByTaskID(seedClientPods, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("0dd6bcd036002bbd79cd8dfebcd9813ec4231562c0b15b02cbc7c70a062717f9").To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("proxy", "download", "range: bytes=100-"), func() {
			clientPod, err := util.ClientExec()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := clientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r 100- -H 'X-Dragonfly-Tag: proxy-bytes-100-' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-proxy-bytes-100-"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "28ded7113a871c1eb08728204810514f08e200952131075ea5f7a3756973ceb2",
				Sha256: "ba9a10ceceb80562fc124dc9bc94ea2a38e3a71e3e746e2140e6381ac791cdeb",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{clientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{clientPod}, util.GetOutputPath("kubectl-proxy-bytes-100-"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(1 * time.Second)
			seedClientPods := make([]*util.PodExec, 3)
			for i := 0; i < 3; i++ {
				seedClientPods[i], err = util.SeedClientExec(i)
				fmt.Println(err)
				Expect(err).NotTo(HaveOccurred())
			}

			sha256sum, err = util.CalculateSha256ByTaskID(seedClientPods, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("proxy", "download", "range: bytes=100-46026751"), func() {
			clientPod, err := util.ClientExec()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := clientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r 100-46026751 -H 'X-Dragonfly-Tag: proxy-bytes-100-46026751' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-proxy-bytes-100-46026751"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "871b9e11903e7f3cadcb25706be7fd2c153b76cdf138f713bd7b506f695fc863",
				Sha256: "ba9a10ceceb80562fc124dc9bc94ea2a38e3a71e3e746e2140e6381ac791cdeb",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{clientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{clientPod}, util.GetOutputPath("kubectl-proxy-bytes-100-46026751"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(1 * time.Second)
			seedClientPods := make([]*util.PodExec, 3)
			for i := 0; i < 3; i++ {
				seedClientPods[i], err = util.SeedClientExec(i)
				fmt.Println(err)
				Expect(err).NotTo(HaveOccurred())
			}

			sha256sum, err = util.CalculateSha256ByTaskID(seedClientPods, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("proxy", "download", "range: bytes=0-46026751"), func() {
			clientPod, err := util.ClientExec()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := clientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r 0-46026751 -H 'X-Dragonfly-Tag: proxy-bytes-0-46026751' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-proxy-bytes-0-46026751"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "28ded7113a871c1eb08728204810514f08e200952131075ea5f7a3756973ceb2",
				Sha256: "327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{clientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{clientPod}, util.GetOutputPath("kubectl-proxy-bytes-0-46026751"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(1 * time.Second)
			seedClientPods := make([]*util.PodExec, 3)
			for i := 0; i < 3; i++ {
				seedClientPods[i], err = util.SeedClientExec(i)
				fmt.Println(err)
				Expect(err).NotTo(HaveOccurred())
			}

			sha256sum, err = util.CalculateSha256ByTaskID(seedClientPods, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("proxy", "download", "range: bytes=100-10240"), func() {
			clientPod, err := util.ClientExec()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := clientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r 100-10240 -H 'X-Dragonfly-Tag: proxy-bytes-100-10240' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-proxy-bytes-100-10240"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "e5ccafcd80ee0dfc48aa1f313ef6086a1fd1e0021183452c096ae40d87d191c6",
				Sha256: "8732360b941ad09a5e0e5d5f9891118bc068f6d0d5a56e3c6d483e4600fbc43f",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{clientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("0dd6bcd036002bbd79cd8dfebcd9813ec4231562c0b15b02cbc7c70a062717f9").To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{clientPod}, util.GetOutputPath("kubectl-proxy-bytes-100-10240"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(1 * time.Second)
			seedClientPods := make([]*util.PodExec, 3)
			for i := 0; i < 3; i++ {
				seedClientPods[i], err = util.SeedClientExec(i)
				fmt.Println(err)
				Expect(err).NotTo(HaveOccurred())
			}

			sha256sum, err = util.CalculateSha256ByTaskID(seedClientPods, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("0dd6bcd036002bbd79cd8dfebcd9813ec4231562c0b15b02cbc7c70a062717f9").To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("proxy", "download", "range: bytes=100-1024"), func() {
			clientPod, err := util.ClientExec()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := clientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r 100-1024 -H 'X-Dragonfly-Tag: proxy-bytes-100-1024' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-proxy-bytes-100-1024"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "16d94182383e84a73f9c85a28a25d781a8a4e56ec37c6e3e61d1f946fb236e85",
				Sha256: "d6d17dca18b8de59e38da525dc24c47b74fec1a790a9f64afdd6538f4f8fa90e",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{clientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("0dd6bcd036002bbd79cd8dfebcd9813ec4231562c0b15b02cbc7c70a062717f9").To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{clientPod}, util.GetOutputPath("kubectl-proxy-bytes-100-1024"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(1 * time.Second)
			seedClientPods := make([]*util.PodExec, 3)
			for i := 0; i < 3; i++ {
				seedClientPods[i], err = util.SeedClientExec(i)
				fmt.Println(err)
				Expect(err).NotTo(HaveOccurred())
			}

			sha256sum, err = util.CalculateSha256ByTaskID(seedClientPods, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("0dd6bcd036002bbd79cd8dfebcd9813ec4231562c0b15b02cbc7c70a062717f9").To(Equal(sha256sum))
		})
	})
})

var _ = Describe("Download Using Prefetch Proxy", func() {
	Context("/etc/containerd/config.toml file", func() {
		It("download should be ok", Label("prefetch-proxy", "download"), func() {
			seedClientPod, err := util.SeedClientExec(0)
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := seedClientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -H 'X-Dragonfly-Tag: prefetch-proxy' %s --output %s", util.GetFileURL("/etc/containerd/config.toml"), util.GetOutputPath("config-prefetch-proxy.coml"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "0fa5f40dfccfcd4c20b2d2dba0bab7205bfced189cc7a013919a6d155d841dab",
				Sha256: "66404431f9a0d5c78205e5a3eb041f76767094fc278c0a091d4ffa10f06cf641",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{seedClientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{seedClientPod}, util.GetOutputPath("config-prefetch-proxy.coml"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file", func() {
		It("download should be ok", Label("prefetch-proxy", "download"), func() {
			seedClientPod, err := util.SeedClientExec(0)
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := seedClientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -H 'X-Dragonfly-Tag: prefetch-proxy' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-prefetch-proxy"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "8cdfeba7c8858b1bf7dc6097aeb6bae12e3a4cb8ed32cc3b60bdb45765bda590",
				Sha256: "327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{seedClientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{seedClientPod}, util.GetOutputPath("kubectl-prefetch-proxy"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))
		})
	})

	Context("/bin/x86_64 file", func() {
		It("download should be ok", Label("prefetch-proxy", "download"), func() {
			seedClientPod, err := util.SeedClientExec(0)
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := seedClientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -H 'X-Dragonfly-Tag: prefetch-proxy' %s --output %s", util.GetFileURL("/bin/x86_64"), util.GetOutputPath("x86_64-prefetch-proxy"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "86b0c4424368f35e500045d5375b8b53940cadf69b71d0ae4b349ac972ad6bdb",
				Sha256: "a1cbf1bf2d66757121677fd7fefafacd4f843a2cb44a451131002803bae56a65",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{seedClientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{seedClientPod}, util.GetOutputPath("x86_64-prefetch-proxy"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))
		})
	})

	Context("/bin/zless file", func() {
		It("download should be ok", Label("prefetch-proxy", "download"), func() {
			seedClientPod, err := util.SeedClientExec(0)
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := seedClientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -H 'X-Dragonfly-Tag: prefetch-proxy' %s --output %s", util.GetFileURL("/bin/zless"), util.GetOutputPath("zless-prefetch-proxy"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "d759035bb589e9334717660701c58ef91ebab49bb5790367ede36bf78daa52a1",
				Sha256: "b0cfe211f851049a78f5812cf5b1d7cb4f3fbb101c02c1865c940d5797f4b73b",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{seedClientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{seedClientPod}, util.GetOutputPath("zless-prefetch-proxy"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))
		})
	})

	Context("/bin/bash file", func() {
		It("download should be ok", Label("prefetch-proxy", "download"), func() {
			seedClientPod, err := util.SeedClientExec(0)
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := seedClientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -H 'X-Dragonfly-Tag: prefetch-proxy' %s --output %s", util.GetFileURL("/bin/bash"), util.GetOutputPath("bash-prefetch-proxy"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "ecd74f0e4f9ff453e3c165934d473ef6946b088fecabb8a2eb7d11d2948f743c",
				Sha256: "c37f93c73cf2f303f874c094f6f76e47b2421a3da9f0e7e0b98bea8a3d685322",
			}

			sha256sum, err := util.CalculateSha256ByTaskID([]*util.PodExec{seedClientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			sha256sum, err = util.CalculateSha256ByOutput([]*util.PodExec{seedClientPod}, util.GetOutputPath("bash-prefetch-proxy"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("prefetch-proxy", "download", "range: bytes=-100"), func() {
			seedClientPod, err := util.SeedClientExec(0)
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := seedClientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r -100 -H 'X-Dragonfly-Tag: prefetch-proxy-bytes-100' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-prefetch-proxy-bytes-100"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "2d9e6ebe90a230ff5ffad7938163456c308e072e910eac3dd063f64f62112f05",
				Sha256: "cd00e292c5970d3c5e2f0ffa5171e555bc46bfc4faddfb4a418b6840b86e79a3",
			}

			sha256sum, err := util.CalculateSha256ByOutput([]*util.PodExec{seedClientPod}, util.GetOutputPath("kubectl-prefetch-proxy-bytes-100"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(5 * time.Second)
			sha256sum, err = util.CalculateSha256ByTaskID([]*util.PodExec{seedClientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))

		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("prefetch-proxy", "download", "range: bytes=0-100"), func() {
			seedClientPod, err := util.SeedClientExec(0)
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := seedClientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r 0-100 -H 'X-Dragonfly-Tag: prefetch-proxy-bytes-0-100' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-prefetch-proxy-bytes-0-100"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "bd969c992e6faff240ac1601afe07bdeaa64c2b75bd5abb36a1fd6601895c542",
				Sha256: "7bbfcb694f6acc69483751d3b48d5bdbdcb284f9e04c7b7caa04b5977cffddc8",
			}

			sha256sum, err := util.CalculateSha256ByOutput([]*util.PodExec{seedClientPod}, util.GetOutputPath("kubectl-prefetch-proxy-bytes-0-100"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(5 * time.Second)
			sha256sum, err = util.CalculateSha256ByTaskID([]*util.PodExec{seedClientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("prefetch-proxy", "download", "range: bytes=100-"), func() {
			seedClientPod, err := util.SeedClientExec(0)
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := seedClientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r 100- -H 'X-Dragonfly-Tag: prefetch-proxy-bytes-100-' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-prefetch-proxy-bytes-100-"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "a58f439c45c5737613f1c1d4f10b57c06239b675fccdc32a3d078af89d2bd870",
				Sha256: "ba9a10ceceb80562fc124dc9bc94ea2a38e3a71e3e746e2140e6381ac791cdeb",
			}

			sha256sum, err := util.CalculateSha256ByOutput([]*util.PodExec{seedClientPod}, util.GetOutputPath("kubectl-prefetch-proxy-bytes-100-"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(5 * time.Second)
			sha256sum, err = util.CalculateSha256ByTaskID([]*util.PodExec{seedClientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("prefetch-proxy", "download", "range: bytes=100-46026751"), func() {
			seedClientPod, err := util.SeedClientExec(0)
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := seedClientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r 100-46026751 -H 'X-Dragonfly-Tag: prefetch-proxy-bytes-100-46026751' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-prefetch-proxy-bytes-100-46026751"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "ba4b4a0a1be91ab6b2a76f99bbd8eac5977c388cfc7de73707d30e5d5cf0a6d4",
				Sha256: "ba9a10ceceb80562fc124dc9bc94ea2a38e3a71e3e746e2140e6381ac791cdeb",
			}

			sha256sum, err := util.CalculateSha256ByOutput([]*util.PodExec{seedClientPod}, util.GetOutputPath("kubectl-prefetch-proxy-bytes-100-46026751"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(5 * time.Second)
			sha256sum, err = util.CalculateSha256ByTaskID([]*util.PodExec{seedClientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("prefetch-proxy", "download", "range: bytes=0-46026751"), func() {
			seedClientPod, err := util.SeedClientExec(0)
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := seedClientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r 0-46026751 -H 'X-Dragonfly-Tag: prefetch-proxy-bytes-0-46026751' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-prefetch-proxy-bytes-0-46026751"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "66b8f945a2906550194aefee5a6da6c3a4c60b533e30340707c8bea67ce35eae",
				Sha256: "327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a",
			}

			sha256sum, err := util.CalculateSha256ByOutput([]*util.PodExec{seedClientPod}, util.GetOutputPath("kubectl-prefetch-proxy-bytes-0-46026751"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(5 * time.Second)
			sha256sum, err = util.CalculateSha256ByTaskID([]*util.PodExec{seedClientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("prefetch-proxy", "download", "range: bytes=100-10240"), func() {
			seedClientPod, err := util.SeedClientExec(0)
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := seedClientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r 100-10240 -H 'X-Dragonfly-Tag: prefetch-proxy-bytes-100-10240' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-prefetch-proxy-bytes-100-10240"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "5cfbab580dd42ef5c3a4d28488fde7e82220115dffc8aff01c0c568665405bcb",
				Sha256: "8732360b941ad09a5e0e5d5f9891118bc068f6d0d5a56e3c6d483e4600fbc43f",
			}

			sha256sum, err := util.CalculateSha256ByOutput([]*util.PodExec{seedClientPod}, util.GetOutputPath("kubectl-prefetch-proxy-bytes-100-10240"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(5 * time.Second)
			sha256sum, err = util.CalculateSha256ByTaskID([]*util.PodExec{seedClientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))
		})
	})

	Context("/bin/kubectl file and set range header", func() {
		It("download should be ok", Label("prefetch-proxy", "download", "range: bytes=100-1024"), func() {
			seedClientPod, err := util.SeedClientExec(0)
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())

			out, err := seedClientPod.Command("sh", "-c", fmt.Sprintf("curl -x 127.0.0.1:4001 -r 100-1024 -H 'X-Dragonfly-Tag: prefetch-proxy-bytes-100-1024' %s --output %s", util.GetFileURL("/bin/kubectl"), util.GetOutputPath("kubectl-prefetch-proxy-bytes-100-1024"))).CombinedOutput()
			fmt.Println(err)
			Expect(err).NotTo(HaveOccurred())
			fmt.Println(string(out))

			fileMetadata := util.FileMetadata{
				ID:     "e376250edb769e967fe48bdd05d50f4124f89f3db367b0eee9850128c93043d7",
				Sha256: "d6d17dca18b8de59e38da525dc24c47b74fec1a790a9f64afdd6538f4f8fa90e",
			}

			sha256sum, err := util.CalculateSha256ByOutput([]*util.PodExec{seedClientPod}, util.GetOutputPath("kubectl-prefetch-proxy-bytes-100-1024"))
			Expect(err).NotTo(HaveOccurred())
			Expect(fileMetadata.Sha256).To(Equal(sha256sum))

			time.Sleep(5 * time.Second)
			sha256sum, err = util.CalculateSha256ByTaskID([]*util.PodExec{seedClientPod}, fileMetadata.ID)
			Expect(err).NotTo(HaveOccurred())
			Expect("327b4022d0bfd1d5e9c0701d4a3f989a536f7e6e865e102dcd77c7e7adb31f9a").To(Equal(sha256sum))
		})
	})
})
