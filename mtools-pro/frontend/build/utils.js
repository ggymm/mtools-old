import path from 'path'

/**
 * * 项目根路径
 * @description 结尾不带/
 */
export function getRootPath() {
  return path.resolve(process.cwd())
}

/**
 * * 项目src路径
 * @description 结尾不带斜杠
 */
export function getSrcPath() {
  return path.resolve(getRootPath(), 'src')
}

/**
 * * 项目wails路径
 * @description 结尾不带斜杠
 */
export function getWailsPath() {
  return path.resolve(getRootPath(), 'wailsjs/go')
}
