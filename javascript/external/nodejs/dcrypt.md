 - 간단한 사용법

```js
const bcryptjs = require('bcryptjs');

const newPasswordHash = bcryptjs.hashSync('123123', 10)

console.log(newPasswordHash)

const result = "$2a$10$TK9R2dmytDIsAKQSi1Sfp.Gbrz5RFzTnbMG6YriMcU7o5tU/2zgm6"

console.log(bcryptjs.compareSync('12323', result))
```