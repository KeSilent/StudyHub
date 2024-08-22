/**
 * 格式化日期
 * @returns
 */
export const formatDate = (timestamp) => {
  const date = new Date(timestamp);
  const year = date.getFullYear();
  const month = ('0' + (date.getMonth() + 1)).slice(-2); // 确保月份是两位数
  const day = ('0' + date.getDate()).slice(-2); // 确保日期是两位数
  return `${year}-${month}-${day}`;
};
