/**
 * @see https://prettier.io/docs/en/configuration.html
 * @type {import("prettier").Config}
 */
const config = {
  tabWidth: 2,
  printWidth: 120,
  singleQuote: true,
  trailingComma: 'all',
  jsxSingleQuote: false,
  bracketSpacing: true,

  plugins: ['@trivago/prettier-plugin-sort-imports'],
  importOrder: [
    '^(react|react-dom/client|react-router)$',
    '^@emotion',
    '^@?(?!repo)[A-Za-z0-9_-]+', // 导入第三方库
    '^@/(?!.*.(svg|png)$).+$',
    '^[../](?!.*\\.(svg|png)$)',
    '^[./](?!.*\\.(svg|png)$)',
    '.*\\.(svg|png)$', // 将 svg 和 png 文件排在最后
  ],
  importOrderSortSpecifiers: true,
  importOrderSideEffects: false,
  importOrderSeparation: false,
};

export default config;
