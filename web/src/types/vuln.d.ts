  // 定义漏洞信息对象的类型
  export interface VulnInfo {
    muna_name: string;
    title: string;
    desc: string;
    created_by: string;
    muna_domain: string;
    cate_type: number;
    type: number;
    attribute: string;
    detail: string;
    repair_suggestion: string;
    url: string;
    poc: string;
    province: string;
    city: string;
    county: string;
    industry: string[];
    attachment_id: string;
    attachment_name: string;
    level: string;
  }
  