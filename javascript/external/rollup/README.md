```json
{
  "devDependencies": {
    "@babel/core": "^7.20.2",
    "@babel/preset-env": "^7.20.2",
    "@rollup/plugin-babel": "^6.0.2",
    "@rollup/plugin-commonjs": "^23.0.2",
    "jest": "^29.3.1",
    "rollup": "^3.3.0"
  }
}

```

```js
import commonjs from '@rollup/plugin-commonjs'
import babel, { getBabelOutputPlugin } from '@rollup/plugin-babel';

export default {
    input: './src/index.js',
    output: [
        {
            file: './build/bundle.es5.js',
            format: 'es',
        }
    ],
    plugins: [
        commonjs(),
        getBabelOutputPlugin({ presets: ['@babel/preset-env'] }),
    ]
}
```

## # 예시

```js
import commonjs from '@rollup/plugin-commonjs'
import { getBabelOutputPlugin } from '@rollup/plugin-babel';
import { nodeResolve } from "@rollup/plugin-node-resolve";
const input = ["src/index.js"];

export default [
    {
        input,
        plugins: [
            getBabelOutputPlugin({ presets: ['@babel/preset-env'] }),
            commonjs(),
        ],
        output: [
            {
                dir: "dist/esm",
                format: "esm",
                // exports: "named",
                // sourcemap: true,
            }
        ]
    },
    {
        input,
        plugins: [
            commonjs(),
            // commonjs({
            //     defaultIsModuleExports: false,
            // }),
            // nodeResolve(),
            // getBabelOutputPlugin({ presets: [['@babel/env']] }),
        ],
        output: [
            {
                file: "dist/type-flow.bundle.js",
                format: "umd",
                name: "TypeFlow"
                // exports: "named",
                // sourcemap: false,
            }
        ]
    },
    {
        input,
        plugins: [
            nodeResolve(), 
            getBabelOutputPlugin({ presets: ['@babel/preset-env'] }),
            commonjs(),
        ],
        output: [
            {
                dir: "dist/cjs",
                format: "cjs",
                // exports: "named",
                // sourcemap: true,
            },
        ],
    },
];
```