const postcssPresetEnv = require('postcss-preset-env');

const yourConfig = {
    autoprefixer: true,
    plugins: [
        postcssPresetEnv({stage: 4})
    ]
}