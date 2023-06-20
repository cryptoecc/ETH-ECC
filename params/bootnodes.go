// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303",   // bootnode-aws-ap-southeast-1-001
	"enode://22a8232c3abc76a16ae9d6c3b164f98775fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de@3.209.45.79:30303",     // bootnode-aws-us-east-1-001
	"enode://8499da03c47d637b20eee24eec3c356c9a2e6148d6fe25ca195c7949ab8ec2c03e3556126b0d7ed644675e78c4318b08691b7b57de10e5f0d40d05b09238fa0a@52.187.207.27:30303",   // bootnode-azure-australiaeast-001
	"enode://103858bdb88756c71f15e9b5e09b56dc1be52f0a5021d46301dbbfb7e130029cc9d0d6f73f693bc29b665770fff7da4d34f3c6379fe12721b5d7a0bcb5ca1fc1@191.234.162.198:30303", // bootnode-azure-brazilsouth-001
	"enode://715171f50508aba88aecd1250af392a45a330af91d7b90701c436b618c86aaa1589c9184561907bebbb56439b8f8787bc01f49a7c77276c58c1b09822d75e8e8@52.231.165.108:30303",  // bootnode-azure-koreasouth-001
	"enode://5d6d7cd20d6da4bb83a1d28cadb5d409b64edf314c0335df658c1a54e32c7c4a7ab7823d57c39b6a757556e68ff1df17c748b698544a55cb488b52479a92b60f@104.42.217.25:30303",   // bootnode-azure-westus-001
	"enode://2b252ab6a1d0f971d9722cb839a42cb81db019ba44c08754628ab4a823487071b5695317c8ccd085219c3a03af063495b2f1da8d18218da2d6a82981b45e6ffc@65.108.70.101:30303",   // bootnode-hetzner-hel
	"enode://4aeb4ab6c14b23e2c4cfdce879c04b0748a20d8e9b59e25ded2a08143e265c6c25936e74cbc8e641e3312ca288673d91f2f93f8e277de3cfa444ecdaaf982052@157.90.35.166:30303",   // bootnode-hetzner-fsn
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// SepoliaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Sepolia test network.
var SepoliaBootnodes = []string{
	// geth
	"enode://9246d00bc8fd1742e5ad2428b80fc4dc45d786283e05ef6edbd9002cbc335d40998444732fbe921cb88e1d2c73d1b1de53bae6a2237996e9bfe14f871baf7066@18.168.182.86:30303",
	// besu
	"enode://ec66ddcf1a974950bd4c782789a7e04f8aa7110a72569b6e65fcd51e937e74eed303b1ea734e4d19cfaec9fbff9b6ee65bf31dcb50ba79acce9dd63a6aca61c7@52.14.151.177:30303",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// LveBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Lve network.
var GwangjuBootnodes = []string{
	"enode://cf69273d0e4514ab18e17c154b0fed66a8a91ebf25b36cced53bcd530cf4cd032c4b14801f692719abe2ebf84d454ee453775bb76ec3c49c440d359a16b24c50@43.200.141.6:30303",
	"enode://922cc8c431841ba4742b8434cd79b76b1952eb8a4ee04dc2c34cb5887d781cb857c39fc2292d0aa08cff7c1ae135ad12d78ec245a35babf4a7af16564ce01a44@35.76.71.80:30301",
	"enode://11d3dafc1773cdf6cea238dc649ea6f9b4b3b793315b71b8f37814510413ab6778dc5b1a67b904336fda3e9886248c7a4312911fea79e35763db44b50d1ed0a6@3.39.197.118:30303",
	"enode://e4f833468b61c913c932e7b28bd7d6e7c9e5ff7210455d217e74fb1fe0a7605e7d860765685257a3ed0204bb52c6d20b6c74e468d9c2ea174ffab82ec6dcfc40@3.39.39.150:30303",
	"enode://c7d8921a939fba36b440ba2bf890013aec7c9236464a616bc812108ed42932f5e15eecc036a5f94912665225b3afb606402331f77b847cb64149b1462518250b@3.36.252.183:30303",
	"enode://11692b68fef3db3feb495a5bfd0796fc3863e726e241dfe452c4835342ccbc8f560a4977ae3d4e794e92d5925540eabdec407f6386239b5dec699d9ab4a8464f@3.1.96.244:30303",
	"enode://adbe0b88a3cb886583ab77e2d7ec96110642bd163f76303e82a7f5355a707d970e906308fde226c2d98ce8057f22bb40adb2f270df357428f78668d25e5366ec@13.215.97.146:30303",
	"enode://b9c1e675cd63f93f81a0a34f813ca5235fb80b7859cf0f38be7aebaffafa2607c174c7210d650a6fc1c0e7388f8a3c2eb4fa2752ffda878bef712991d4f459c0@54.90.67.185:30303",   // bootnode-US EAST(N. Virginia)
	"enode://90913f13fb511de0a239ff09dc6b5285bf64246bee591586fb6ab378f777db57751e6c3dc68b5877c60b558bec566f4769ebaa77390123aa48547ea08b08dd80@18.191.247.176:30303", // bootnode-US EAST(Ohio)
	"enode://6f20b438f270e58c127cac9be2fb207c4cf46268d77dc441d1f60cdb7dacc9a77d1a32382e3b7fc3e2528b03e844ad05781d1edf4e0cae1d2553c071ae224cfa@54.193.233.245:30303", // bootnode-US WEST(N. California)
	"enode://195ce0bebda69801ed2ed7e155e2c94a3332dd8266bd01f8802b9f56b4fbfe7c757cefaae47e83ce02955147c9f9c71b3a5037eba740421e30841ee8e0b7083d@35.88.132.17:30303",   // bootnode-US WEST(Oregon)
	"enode://fce6be61c4dc6f0cd363bb39382e42e8d27bb482af47e0f4a2e3036428e8a99d12afe8db900c668ad2be154e0811a391263c082e2e68289ad39503a64d8361cc@3.75.87.235:30303",    // bootnode-EU Frankfurt
	"enode://a3195014477e7035476ce06f763f9f8c949fecdccbe91acd39dca018600ea57d2547607cbd7751ba0e8af10bf89717b5c1cb8ca5704abc27b396ec77c4591041@3.250.170.177:30303",  // bootnode-EU Ireland
	"enode://60d24b1122f8c3233aa44e18b871a92991f006c463bd199592728eab16ffb28cccd626a80369a59c486195ac126747eb4043791d3fc9aa3b9f979aaafef26b98@18.170.212.3:30303",   // bootnode-EU London
	"enode://4f0f2a5d723f46b93bf258eff8febfc88a20ccd007218a168b53057852446ccb4614c7e53549163faa1212f2a211dcd42aa9d060c6bd7243b470ca7db73b52d8@35.180.65.141:30303",  // bootnode-EU Paris
	"enode://df4d94020b1f9b024aa2f0e6bf785b3895290247c8a34d3b1b7f5cc626d70d1a7112c52b650794319963247ca2547acd97adbc23dc5b7c27cdbcc8adb5dfe6e3@13.48.248.67:30303",   // bootnode-EU Stockholm
}

// SeoulBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Seoul network.
var SeoulBootnodes = []string{
	"enode://205624a79f175a7c1332a1bf11b94691dd996f876679dbcb1258bd28750b4d1d91d4deaa3567e55557b8b7503c05984733654c3badfd270e46b8980db1e64328@43.200.141.6:30303",
}

var WorldlandtestBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://3f87316c80ee6ec015e21ce7e2cd5b1067d946a52e87ae83aacc03fa050379c1c3c5412fc31a7092e210636848e2d57db28453779aa1ce7da7113bcd80060adf@65.109.62.117:30303",
	"enode://5c771ff0c5a0400edaeb12dfb9b405b35a3b6fcb1e780cc223eebb050563f67e122fbbbd0d617f89c2163f97062a9bc89b8ba4c03d039b2d09ba7ce8d2c3c979@35.82.33.142:30303",
	"enode://8a66b37212a72f338e86c63cc05fa7a1faa954bd384a48f3656032239388fa889c9b59b01879078a44776ca0c4562137a929d3a78662cfc3de36af154937f81d@176.9.38.209:30303",
	"enode://aae79ff5076aecaef5e1748bb50f02bb14703823200f184c7beb9b57b12ebe58416dfc96dd41a459fcd91c7931f9c05323b124be03ee2f5a19febaee0d15678a@65.109.62.118:30303",
	"enode://873631a6313fe26e309b566fa16b8b12b9e37bf019180e7c18e982290cc3df489664d57480e52a0acb6f8c33322010379bcf11173bcfb018a91f06b475c4bac1@144.76.17.180:30303",
	"enode://efe4155a02f678d8f85a66e9d0c6aac7e7f68e94735aaca4410f8f1ed5b04c8134e619b98ca2b81d4debae76bb36074c43bfcf278bfad1babb90218784865801@54.176.239.225:30303",
	"enode://a2cb0517c88b71580ab46595c5dab4aca73d248404b7d1773bb089a1ea5a0cbd88b2c6a315ff3199a62c804b05783c568628933facf351f6e7f88eef527f51ea@54.177.117.107:30303",
	"enode://8cf6f46d351c319c708e54eb693050ef8e717115c96073d6961e2ded71fc609f3b2f8ba42ec005ab554e4d4147393fe12c1bf7ce88b6353585897e9c758bb3a2@54.156.57.38:30303",
	"enode://8b05b73b65e6307a09e0c7a92d2b8a83384939f81618783069de894a9ba8bde42ea5a88f52c02d54fc9ffa12b080fb5cf625412a2a63d0f9baa19dae9f7cce39@35.158.233.108:30303",
	"enode://c6bdbf1d7eb319f6ea54df20a270fb6bc15d31f7b8b9dcaf918f8161277769038df7834d0dc94b250b1e92017b477fd2c19ce920c7a30b1dbb9eea3769108c61@3.38.46.69:30303",
	"enode://dafb0b584ba443219fea0bd172e9d3457eabd8c913e7aed7ab9627577850355903b85c64611de056433458d89a4e39c5d48fb44f18181d650fedf43808368d70@44.230.118.175:30303",
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://b5948a2d3e9d486c4d75bf32713221c2bd6cf86463302339299bd227dc2e276cd5a1c7ca4f43a0e9122fe9af884efed563bd2a1fd28661f3b5f5ad7bf1de5949@18.218.250.66:30303",

	// Ethereum Foundation bootnode
	"enode://a61215641fb8714a373c80edbfa0ea8878243193f57c96eeb44d0bc019ef295abd4e044fd619bfc4c59731a73fb79afe84e9ab6da0c743ceb479cbb6d263fa91@3.11.147.67:30303",

	// Goerli Initiative bootnodes
	"enode://d4f764a48ec2a8ecf883735776fdefe0a3949eb0ca476bd7bc8d0954a9defe8fea15ae5da7d40b5d2d59ce9524a99daedadf6da6283fca492cc80b53689fb3b3@46.4.99.122:32109",
	"enode://d2b720352e8216c9efc470091aa91ddafc53e222b32780f505c817ceef69e01d5b0b0797b69db254c586f493872352f5a022b4d8479a00fc92ec55f9ad46a27e@88.99.70.182:30303",
}

var KilnBootnodes = []string{
	"enode://c354db99124f0faf677ff0e75c3cbbd568b2febc186af664e0c51ac435609badedc67a18a63adb64dacc1780a28dcefebfc29b83fd1a3f4aa3c0eb161364cf94@164.92.130.5:30303",
	"enode://d41af1662434cad0a88fe3c7c92375ec5719f4516ab6d8cb9695e0e2e815382c767038e72c224e04040885157da47422f756c040a9072676c6e35c5b1a383cce@138.68.66.103:30303",
	"enode://91a745c3fb069f6b99cad10b75c463d527711b106b622756e9ef9f12d2631b6cb885f831d1c8731b9bc7177cae5e1ea1f1be087f86d7d30b590a91f22bc041b0@165.232.180.230:30303",
	"enode://b74bd2e8a9f0c53f0c93bcce80818f2f19439fd807af5c7fbc3efb10130c6ee08be8f3aaec7dc0a057ad7b2a809c8f34dc62431e9b6954b07a6548cc59867884@164.92.140.200:30303",
}

var V5Bootnodes = []string{
	// Teku team's bootnode
	"enr:-KG4QOtcP9X1FbIMOe17QNMKqDxCpm14jcX5tiOE4_TyMrFqbmhPZHK_ZPG2Gxb1GE2xdtodOfx9-cgvNtxnRyHEmC0ghGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQDE8KdiXNlY3AyNTZrMaEDhpehBDbZjM_L9ek699Y7vhUJ-eAdMyQW_Fil522Y0fODdGNwgiMog3VkcIIjKA",
	"enr:-KG4QDyytgmE4f7AnvW-ZaUOIi9i79qX4JwjRAiXBZCU65wOfBu-3Nb5I7b_Rmg3KCOcZM_C3y5pg7EBU5XGrcLTduQEhGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQ2_DUbiXNlY3AyNTZrMaEDKnz_-ps3UUOfHWVYaskI5kWYO_vtYMGYCQRAR3gHDouDdGNwgiMog3VkcIIjKA",
	// Prylab team's bootnodes
	"enr:-Ku4QImhMc1z8yCiNJ1TyUxdcfNucje3BGwEHzodEZUan8PherEo4sF7pPHPSIB1NNuSg5fZy7qFsjmUKs2ea1Whi0EBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQOVphkDqal4QzPMksc5wnpuC3gvSC8AfbFOnZY_On34wIN1ZHCCIyg",
	"enr:-Ku4QP2xDnEtUXIjzJ_DhlCRN9SN99RYQPJL92TMlSv7U5C1YnYLjwOQHgZIUXw6c-BvRg2Yc2QsZxxoS_pPRVe0yK8Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMeFF5GrS7UZpAH2Ly84aLK-TyvH-dRo0JM1i8yygH50YN1ZHCCJxA",
	"enr:-Ku4QPp9z1W4tAO8Ber_NQierYaOStqhDqQdOPY3bB3jDgkjcbk6YrEnVYIiCBbTxuar3CzS528d2iE7TdJsrL-dEKoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMw5fqqkw2hHC4F5HZZDPsNmPdB1Gi8JPQK7pRc9XHh-oN1ZHCCKvg",
	// Lighthouse team's bootnodes
	"enr:-IS4QLkKqDMy_ExrpOEWa59NiClemOnor-krjp4qoeZwIw2QduPC-q7Kz4u1IOWf3DDbdxqQIgC4fejavBOuUPy-HE4BgmlkgnY0gmlwhCLzAHqJc2VjcDI1NmsxoQLQSJfEAHZApkm5edTCZ_4qps_1k_ub2CxHFxi-gr2JMIN1ZHCCIyg",
	"enr:-IS4QDAyibHCzYZmIYZCjXwU9BqpotWmv2BsFlIq1V31BwDDMJPFEbox1ijT5c2Ou3kvieOKejxuaCqIcjxBjJ_3j_cBgmlkgnY0gmlwhAMaHiCJc2VjcDI1NmsxoQJIdpj_foZ02MXz4It8xKD7yUHTBx7lVFn3oeRP21KRV4N1ZHCCIyg",
	// EF bootnodes
	"enr:-Ku4QHqVeJ8PPICcWk1vSn_XcSkjOkNiTg6Fmii5j6vUQgvzMc9L1goFnLKgXqBJspJjIsB91LTOleFmyWWrFVATGngBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAMRHkWJc2VjcDI1NmsxoQKLVXFOhp2uX6jeT0DvvDpPcU8FWMjQdR4wMuORMhpX24N1ZHCCIyg",
	"enr:-Ku4QG-2_Md3sZIAUebGYT6g0SMskIml77l6yR-M_JXc-UdNHCmHQeOiMLbylPejyJsdAPsTHJyjJB2sYGDLe0dn8uYBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhBLY-NyJc2VjcDI1NmsxoQORcM6e19T1T9gi7jxEZjk_sjVLGFscUNqAY9obgZaxbIN1ZHCCIyg",
	"enr:-Ku4QPn5eVhcoF1opaFEvg1b6JNFD2rqVkHQ8HApOKK61OIcIXD127bKWgAtbwI7pnxx6cDyk_nI88TrZKQaGMZj0q0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg",
	"enr:-Ku4QEWzdnVtXc2Q0ZVigfCGggOVB2Vc1ZCPEc6j21NIFLODSJbvNaef1g4PxhPwl_3kax86YPheFUSLXPRs98vvYsoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDZBrP2Jc2VjcDI1NmsxoQM6jr8Rb1ktLEsVcKAPa08wCsKUmvoQ8khiOl_SLozf9IN1ZHCCIyg",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	case SepoliaGenesisHash:
		net = "sepolia"
	case GwangjuGenesisHash:
		net = "gwangju"
	case SeoulGenesisHash:
		net = "Seoul"
	case WorldlandtestGenesisHash:
		net = "wordlalndtest"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
