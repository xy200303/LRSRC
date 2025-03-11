export function buildRouteTree(routes) {
    const map = new Map();
    const tree = [];
    // 首先将所有节点存入映射
    routes.forEach(route => {
        map.set(route.path, { ...route, children: [] });
    });
    // 构建树结构
    routes.forEach(route => {
        const pathParts = route.path.split('/').filter(part => part.length > 0);
        let currentLevel = tree;
        for (let i = 0; i < pathParts.length; i++) {
            const part = pathParts[i];
            const fullPath = '/' + pathParts.slice(0, i + 1).join('/');
            let node = currentLevel.find(node => node.path === fullPath);
            if (!node) {
                node = map.get(fullPath);
                if (node) {
                    currentLevel.push(node);
                }
            }
            if (node) {
                currentLevel = node.children;
            }
        }
    });
    return tree;
}