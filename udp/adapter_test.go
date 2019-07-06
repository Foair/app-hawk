package udp

import (
	"encoding/hex"
	"testing"
)

func TestAdapter(T *testing.T) {
	data, _ := hex.DecodeString("4e535543563100034e802219c7ad41557b016ede1f89dbba859dd47d3972711363e292039778c84f14149f0826b58d1f7416e3bb7e2b56d639201d2164963ef6bbe30360aa7ab595eb3bf31c8f94c9272e8be18b47621ee01defa7c97a9e6fb8e0ef8e4f4e37bae2e5cce8704446d35c44c5fa501cf1330481030bed7b7e10bd78536c8835e166c437f62683e4cf3f674d7e5c330ff858c4fb9ef70a8c1e03e8e626f474b966b3cbdc4f717d7b185f36883143acf2de20701e7db94170d6117ceec6f2264089108ea0b2e9368a1cc703507ea5bf09b585cd8724725b0f7535539468cd23e39388c3ef3c443195da98ba6cef1277d1cbc4770e61813ac723f2c1522bf0cfeb93a9b86d0ac0855162bcdced0b6bb70f76cdc56b3652bb906d3f31b236d9b3d97b665229ed119502a91608c887dd747f317fd084c082069429425e386a8754a8e6f675f9a9721f20f238ee9ac91fbdc9d0590a4f1ca662c1cbccd69c3637c8261cb6e89e595ddca86337c1e5903ed30ff1606170a912b5b3c3b9363203b4d752b124bf91e70218c76e7ce0e8303a0414323b9d80c7382086506d3f679738cf1de4e2972590a7c548eb3807c8ed4b41fdf1093087416a9298ada2a83c3da793661638e9f4f2e683c0306a60dfe9aa43b4a268c1f53676d181fa0b64f1b7a632a1f9fe1442ed301b547abc6b7b2cd1c0a626d94b1f4bbaa9e43f60c20063e0626a066f4e27bfaa022ee6f85e730d5616a469a1e5678721b0a33b940d4cb7ed8c482da8a04e87e94b063d234ad566c92914b9f7cbcfe5ccd379cc586b0967c2a98e9099545a86840f90d7af4591094e0790dc53fb2d603c8964de610459e43fefa683015fcaadf9056a08a58c00d9f5c89984f8394e08a480c4064e65924a490a1817681c692e708d86b20b3d3a0bfcacdac5e8cdc700f69cdc1def73a62ddde3a88b6cbf6d1b35c301f7488acb7d133ea644c7b02f6a0e5e8258cedc71fdfef636257f7a7139765f27c3fffbb52b4e0be866836cf5d66a457549c60b519bd4d3efa7c3e9ca4808b3f97d8f803d4adbaa3439d81477f2f177a3f6c2a328cc977cb5dc3bda")
	queue <- data
}
