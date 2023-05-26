// config-overrides.js

const path = require('path');
const fs = require('fs-extra');

module.exports = function override(config) {
  // 修改输出目录为 dist
  config.output.path = path.resolve(__dirname, '../webview/assets');

  // 复制静态文件到自定义输出目录
  const sourceDir = path.resolve(__dirname, 'public');
  const targetDir = path.resolve(__dirname, '../webview/assets');
  fs.copySync(sourceDir, targetDir);

  return config;
};
