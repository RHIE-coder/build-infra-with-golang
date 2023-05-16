# JEST


## # Configuration Example

### - `jest.config.js`

```js
module.exports = {
    testMatch: [
        "<rootDir>/**/*.test.(js|ts)",
        "<rootDir>/(**/test/**/*.spec.(js|ts)|**/__test__/*.(js|ts))",
    ],
    testPathIgnorePatterns: [
        "<rootDir>/node_modules/",
    ],
}
```