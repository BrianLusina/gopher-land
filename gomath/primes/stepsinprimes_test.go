package primes

import (
	"math"
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func isPrimeOPI(n int) bool {
	if n == 2 {
		return true
	}
	if (n < 3) || (n%2 == 0) {
		return false
	}
	limit := int(math.Sqrt(float64(n)) + 1.0)
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func stepOPI(g, m, n int) []int {
	var res []int
	for i := m; i <= n-g; i++ {
		if isPrimeOPI(i) && isPrimeOPI(i+g) {
			res = append(res, i)
			res = append(res, i+g)
			return res
		}
	}
	return res
}

func dotest(k, start, nd int, exp []int) {
	var ans = Step(k, start, nd)
	Expect(ans).To(Equal(exp))
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
func randomdotest() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 25; i++ {
		n := random(1000, 1000000)
		sol := stepOPI(4, n, n+1000)
		dotest(4, n, n+1000, sol)
	}
}

var _ = Describe("Test Example", func() {

	It("should handle basic cases", func() {
		dotest(2, 100, 110, []int{101, 103})
		dotest(4, 100, 110, []int{103, 107})
		dotest(6, 100, 110, []int{101, 107})
		dotest(8, 300, 400, []int{359, 367})
		dotest(10, 300, 400, []int{307, 317})

		dotest(4, 30000, 100000, []int{30109, 30113})
		dotest(6, 30000, 100000, []int{30091, 30097})
		dotest(8, 30000, 100000, []int{30089, 30097})
		dotest(11, 30000, 100000, nil)
		dotest(2, 10000000, 11000000, []int{10000139, 10000141})

		dotest(52, 1300, 15000, []int{1321, 1373})
		dotest(10, 4900, 5000, []int{4909, 4919})
		dotest(30, 4900, 5000, []int{4903, 4933})
		dotest(2, 4900, 5000, []int{4931, 4933})
		dotest(2, 104000, 105000, []int{104087, 104089})

		dotest(2, 4900, 4919, nil)
		dotest(7, 4900, 4919, nil)
		dotest(4, 30115, 100000, []int{30133, 30137})
		dotest(4, 30140, 100000, []int{30319, 30323})
		dotest(4, 30000, 30325, []int{30109, 30113})

		dotest(4, 945104, 995104, []int{945289, 945293})
		dotest(4, 370207, 420207, []int{370213, 370217})
		dotest(4, 916320, 966320, []int{916507, 916511})
		dotest(4, 408393, 458393, []int{408427, 408431})
		dotest(4, 739081, 789081, []int{739099, 739103})
		dotest(4, 526070, 576070, []int{526117, 526121})
		dotest(4, 113907, 163907, []int{114073, 114077})
		dotest(4, 775045, 825045, []int{775087, 775091})
		dotest(4, 124610, 174610, []int{124669, 124673})
		dotest(4, 576021, 626021, []int{576217, 576221})
		dotest(7, 707088, 757088, nil)
		dotest(7, 400955, 450955, nil)

		dotest(2, 664239, 714239, []int{664271, 664273})
		dotest(2, 628311, 678311, []int{628679, 628681})
		dotest(2, 989871, 1039871, []int{989999, 990001})
		dotest(2, 822238, 872238, []int{822389, 822391})
		dotest(2, 182199, 232199, []int{182339, 182341})
		dotest(2, 113554, 163554, []int{113621, 113623})
		dotest(2, 723661, 773661, []int{723797, 723799})
		dotest(2, 89949, 139949, []int{90017, 90019})
		dotest(2, 125810, 175810, []int{125897, 125899})
		dotest(2, 260997, 310997, []int{261011, 261013})
		dotest(2, 711422, 761422, []int{711497, 711499})
		dotest(2, 466484, 516484, []int{466649, 466651})
		dotest(2, 644577, 694577, []int{644597, 644599})
		dotest(2, 483344, 533344, []int{483407, 483409})
		dotest(2, 674498, 724498, []int{674699, 674701})
		dotest(2, 358208, 408208, []int{358277, 358279})
		dotest(2, 915256, 965256, []int{915587, 915589})
		dotest(2, 236708, 286708, []int{236771, 236773})
		dotest(2, 63517, 113517, []int{63587, 63589})
		dotest(2, 133076, 183076, []int{133277, 133279})
		dotest(2, 359079, 409079, []int{359207, 359209})
		dotest(2, 465986, 515986, []int{466181, 466183})
		dotest(2, 331048, 381048, []int{331337, 331339})
		dotest(2, 587771, 637771, []int{587771, 587773})
		dotest(2, 544563, 594563, []int{544721, 544723})
		dotest(2, 603281, 653281, []int{603389, 603391})
		dotest(2, 739840, 789840, []int{739859, 739861})
		dotest(2, 312653, 362653, []int{312677, 312679})
		dotest(2, 192494, 242494, []int{192497, 192499})
		dotest(2, 935263, 985263, []int{935591, 935593})
		dotest(2, 248019, 298019, []int{248117, 248119})
		dotest(2, 334147, 384147, []int{334331, 334333})
		dotest(2, 57508, 107508, []int{57527, 57529})
		dotest(2, 659866, 709866, []int{659999, 660001})
		dotest(2, 279462, 329462, []int{279479, 279481})
		dotest(2, 797154, 847154, []int{797309, 797311})
		dotest(2, 885385, 935385, []int{885551, 885553})
		dotest(2, 699030, 749030, []int{699287, 699289})
		dotest(2, 463755, 513755, []int{463829, 463831})
		dotest(2, 741245, 791245, []int{741341, 741343})
	})
	It("should handle random cases", func() {
		randomdotest()
	})
})
