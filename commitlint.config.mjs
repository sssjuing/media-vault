/** @type {import('cz-git').UserConfig} */
export default {
  extends: ['@commitlint/config-conventional'],
  prompt: {
    useEmoji: true,
    scopes: [{ name: 'backend:   后端应用', value: 'backend' }],
  },
};
