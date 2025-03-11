<template>
    <div>
        <!-- 发表评论 -->
        <div v-clickoutside="hideReplyBtn" @click="inputFocus" class="my-reply">
            <simple-editor height="200px" v-model="replyComment"></simple-editor>
            <div class="reply-btn-box">
                <el-button class="reply-btn" size="medium" @click="sendComment" type="primary">发表评论</el-button>
            </div>
        </div>
        <!-- 评论选项 -->
        <el-tabs v-model="activeName" @tab-click="handleTagClick">
            <el-tab-pane label="最新评论" name="new"></el-tab-pane>
            <el-tab-pane label="最热评论" name="hot"></el-tab-pane>
        </el-tabs>
        <!-- 评论展示部分 -->
        <div v-for="(item,i) in modelValue" :key="i" class="author-title reply-father">
            <el-avatar class="header-img" :size="40" :src="item.avatar"></el-avatar>
            <!-- 作者信息 -->
            <div class="author-info">
                <span class="author-name">{{item.created_by}}
                  <el-button 
                        v-if="item.show_delete" 
                        type="info" 
                        size="small" 
                        link
                        @click="handleDelete(item)"
                    >删除</el-button>
                </span>
                <span class="author-time">{{formatDateChinese(item.created_at)}}</span>
            </div>
            <!-- 评论展示 -->
            <div class="talk-box">
                <div class="reply" v-html="item.content"></div>
            </div>
        </div>
    </div>
</template>
<script setup>
import { ref, onMounted, onUnmounted,watch } from 'vue';
import { ElMessage } from 'element-plus';
import { ChatDotSquare, CaretTop } from '@element-plus/icons-vue';
import { dateStr } from '@/utils';
import {formatDateChinese} from "@/utils/datetime"
import {extractTextFromHtml} from "@/utils"

const replyComment = ref('');
const activeName = ref('new');
const emit = defineEmits(['update:modelValue', 'sendComment', 'tabChange',"deleteComment"]);
const props = defineProps({
  modelValue: {
    type: Array,
    default: ()=>[],
  },
});


const localComments=ref(props.value)
watch(() => props.modelValue, (newVal, oldVal) => {
  localComments.value=newVal
}, { deep: true });

const handleTagClick = () => {
    emit('tabChange', activeName.value);
};

const sendComment = () => {
    if (!extractTextFromHtml(replyComment.value)) {
        ElMessage({
            showClose: true,
            type: 'warning',
            message: '评论不能为空',
        });
    } else {
        //发表评论
        emit("sendComment",replyComment.value)
        replyComment.value=""
    }
};

const handleDelete = (comment) => {
    emit('deleteComment', comment.id);
};
</script>
<style scoped>
.my-reply {
  padding: 10px;
}

.my-reply .reply-btn-box {
  margin-top: 10px;
  text-align: right;
}

.author-title {
  padding: 10px;
  border-bottom: 1px solid rgba(178, 186, 194, 0.3);
}

.author-title .header-img {
  display: inline-block;
  vertical-align: top;
}

.author-title .author-info {
  display: inline-block;
  margin-left: 10px;
  width: 80%;
}

.author-title .author-info > span {
  display: block;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.author-title .author-info .author-name {
  color: #000;
  font-size: 18px;
  font-weight: bold;
}

.author-title .author-info .author-time {
  font-size: 14px;
}

.author-title .talk-box {
  margin-left: 50px;
}

.author-title .talk-box .reply {
  font-size: 16px;
  color: #000;
}

.author-title .talk-box .reply :deep(img) {
  max-width: 200px;
  height: auto;
  object-fit: contain;
}
</style>