/** @type {import('cz-git').UserConfig} */
export default {
  extends: ['@commitlint/config-conventional'],
  prompt: {
    useEmoji: true,
    scopes: [{ value: 'backend', name: 'Backend' }],
  },
};
