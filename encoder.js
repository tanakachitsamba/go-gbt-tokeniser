const { encode, decode } = require('gpt-3-encoder')

const str = process.argv[2];
const encoded = encode(str);

console.log(JSON.stringify({tokens: encoded, count: encoded.length}));
