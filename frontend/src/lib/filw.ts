/**
 * @author 成雨
 * @date 2021/1/28
 * @Description:
 */

/**
 * 截取文件扩展名
 * @param {string} filename
 * @return {string}
 */
export function fExt(filename: string) {
  return filename.split(".").pop();
}

/**
 * 获取文件名 不带有扩展名
 * @param {string} filename
 * @returns {string}
 */
export function fName(filename: string) {
  filename = filename.replace(/\\/g, "/");
  return filename.substring(
    filename.lastIndexOf("/") + 1,
    filename.lastIndexOf(".")
  );
}
