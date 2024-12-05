import dayjs from "dayjs";

// 格式化时间
export function formatTime(time, format = "YYYY-MM-DD HH:mm:ss") {
  return dayjs(time).format(format);
}