// 用于处理CSV文件的工具函数
import Papa from "papaparse";
export function useCsv() {
    // 导出CSV文件
    const exportCsv = (data, fileName = "导出列表") => {
        const csv = Papa.unparse(data);
        const blob = new Blob([csv], { type: "text/csv;charset=utf-8;" });
        const url = URL.createObjectURL(blob);

        const link = document.createElement("a");
        link.setAttribute("href", url);
        link.setAttribute("download", fileName+".csv");
        link.click();
    };
    return {
        exportCsv
    };
}
