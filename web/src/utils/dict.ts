import { ref, toRefs } from "vue";
import useDictStore from "@/stores/dictStores";
import { getDictData } from "@/api/admin/dict";


/**
 * 获取字典数据
 */
export function useDict(...args: string[]) {
    const res = ref<{ [key: string]: any }>({});
    return (() => {
      args.forEach((dictType, index) => {
        res.value[dictType] = [];
        const dicts = useDictStore().getDict(dictType);
        if (dicts) {
          res.value[dictType] = dicts;
        } else {
          getDictData({
            dict_type:dictType
          }).then((resp: { data: any[]; }) => {
            res.value[dictType] = resp.data.map(p => ({ label: p.label_name, value: p.value,el_tag_type:p.el_tag_type,el_tag_effect:p.el_tag_effect }))
            useDictStore().setDict(dictType, res.value[dictType]);
          })
        }
      })
      return toRefs(res.value);
    })()
  }