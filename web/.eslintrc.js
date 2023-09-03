module.exports = {
    root: true,
    plugins: ['@typescript-eslint', 'import', 'prettier'],
    extends: [
        'airbnb-typescript',
        'prettier',
        'plugin:@typescript-eslint/recommended',
        'plugin:import/typescript',
    ],
    parser: '@typescript-eslint/parser',
    parserOptions: {
        project: './tsconfig.json',
        tsconfigRootDir: __dirname,
    },
    rules: {
        "react/jsx-filename-extension": 'off'
    }
};