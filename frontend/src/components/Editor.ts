/**
 * @author 成雨
 * @date 2021/1/30
 * @Description:
 */

export interface CFile {
  id: string;
  filename: string;
  isConverted: boolean;
  convertedPath: string;
  name: string;
  size: number;
  convertedSize: 0;
  isProcessed: boolean;
}
