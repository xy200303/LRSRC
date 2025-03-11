export function dateStr(date){
    let time = new Date().getTime();
    time = parseInt((time - date) / 1000);
    let s;
    if (time < 60 * 10) {
        return '刚刚';
    } else if ((time < 60 * 60) && (time >= 60 * 10)) {
        s = Math.floor(time / 60);
        return `${s}分钟前`;
    } else if ((time < 60 * 60 * 24) && (time >= 60 * 60)) {
        s = Math.floor(time / 60 / 60);
        return `${s}小时前`;
    } else if ((time < 60 * 60 * 24 * 30) && (time >= 60 * 60 * 24)) {
        s = Math.floor(time / 60 / 60 / 24);
        return `${s}天前`;
    } else {
        const dateObj = new Date(parseInt(date));
        return `${dateObj.getFullYear()}/${dateObj.getMonth() + 1}/${dateObj.getDate()}`;
    }
}

export function extractTextFromHtml(htmlString) {
    let parser = new DOMParser();
    let doc = parser.parseFromString(htmlString, 'text/html');
    // 移除所有的 <script> and <style> 元素
    let scriptsAndStyles = doc.querySelectorAll('script, style');
    scriptsAndStyles.forEach(element => element.remove());
    return doc.body.textContent || '';
}


export * from "./auth.ts";
export * from "./common.ts";
export * from "./datetime.ts";
export * from "./requests.ts";
export * from "./file.ts";