<template>
  <div class="aside-container">
    <div class="menu_content">
      <div class="title">
        短链分组
      </div>

      <a-button class="createGroup" type="link" @click="showModal">
        <PlusSquareOutlined />
      </a-button>
    </div>
    <a-menu
      class="groupMenu"
      style="width: 256px"
      :default-selected-keys="[selectedGroup?.gid]"
      mode="inline"
    >
      <a-menu-item 
        v-for="group in groupList" 
        :key="group.gid" 
        @click="handleGroupClick(group)"
        @dblclick="startEdit(group)"
        draggable="true"
        @dragstart="handleDragStart($event, group)"
        @dragend="handleDragEnd"
      >
        <div class="group-item">
          <template v-if="editingGroup?.gid === group.gid">
            <a-input
              v-model:value="editingName"
              size="small"
              style="width: 120px"
              @click.stop
              ref="editInput"
            />
            <div class="edit-actions">
              <CheckOutlined class="edit-icon" @click.stop="confirmEdit" />
              <CloseOutlined class="edit-icon" @click.stop="cancelEdit" />
            </div>
          </template>
          <template v-else>
            {{ group.name }}
          </template>
        </div>
      </a-menu-item>
    </a-menu>

    <!-- 垃圾桶区域 -->
    <div 
      class="trash-zone"
      :class="{ 'trash-active': isDragging }"
      @dragover="handleDragOver"
      @drop="handleDrop"
    >
      <DeleteOutlined />
    </div>

    <!-- 添加创建分组的模态框 -->
    <a-modal
      v-model:visible="modalVisible"
      title="创建分组"
      @ok="handleCreateGroup"
      @cancel="modalVisible = false"
      :confirmLoading="confirmLoading"
    >
      <a-input v-model:value="groupName" placeholder="请输入分组名称" />
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import {
  PlusSquareOutlined, 
  CheckOutlined, 
  CloseOutlined,
  DeleteOutlined
} from "@ant-design/icons-vue";
import { ref, onMounted } from "vue";
import { getGroupInfo, createGroup, updateGroup, deleteGroup } from "@/api/group/groupApi";
import { fetchShortLinks } from "@/api/shortlink/shortlinkApi";
import { message } from 'ant-design-vue';

interface GroupData {
  gid: string;
  name: string;
  username: string;
}

const groupList = ref<GroupData[]>([]);
const selectedGroup = ref<GroupData | null>(null);
const shortLinks = ref([]);

// 新增的响应式变量
const modalVisible = ref(false);
const confirmLoading = ref(false);
const groupName = ref('');

// 新增编辑相关的响应式变量
const editingGroup = ref<GroupData | null>(null);
const editingName = ref('');
const editInput = ref<HTMLElement | null>(null);

// 拖拽相关的响应式变量
const isDragging = ref(false);
const draggingGroup = ref<GroupData | null>(null);

// 获取分组信息并初始化
const initializeGroups = async () => {
  try {
    const username = localStorage.getItem('username');
    if (username) {
      const response = await getGroupInfo(username);
      if (response.data.success) {
        if (response.data.data.data === null) {
          // 如果没有分组数据，自动创建默认分组
          await createDefaultGroup(username);
        } else {
          groupList.value = response.data.data.data;
          selectedGroup.value = groupList.value[0]; // 选定第一个分组
          await fetchShortLinksForGroup(selectedGroup.value.gid);
        }
      }
    }
  } catch (error) {
    console.error("获取群组信息失败:", error);
    message.error('获取群组信息失败');
  }
};

// 创建默认分组
const createDefaultGroup = async (username: string) => {
  try {
    const response = await createGroup({
      name: "默认分组",
      username: username
    });

    if (response.data.success) {
      message.success('已自动创建默认分组');
      // 重新获取分组列表
      await initializeGroups();
    } else {
      message.error(response.data.message || '创建默认分组失败');
    }
  } catch (error) {
    console.error("创建默认分组失败:", error);
    message.error('创建默认分组失败');
  }
};

// 获取短链接信息
const fetchShortLinksForGroup = async (gid: string) => {
  try {
    const response = await fetchShortLinks(gid, 1, 10, 1); // 假设分页参数为 page=1, size=10, orderTag=1
    if (response.data.success) {
      shortLinks.value = response.data.data.list.map(link => ({
        key: link.short_uri,
        info: link.describe,
        date: link.update_at,
        url1: link.full_short_url,
        url2: link.origin_url,
        visitsToday: link.today_pv,
        visitsTotal: link.total_pv,
        visitorsToday: link.today_uv,
        visitorsTotal: link.total_uv,
        ipToday: link.today_uip,
        ipTotal: link.total_uip,
      }));
    } else {
      message.error(response.data.message || '获取短链接失败');
    }
  } catch (error) {
    console.error("获取短链接失败:", error);
    message.error('获取短链接失败');
  }
};

// 处理分组点击
const handleGroupClick = async (group: GroupData) => {
  selectedGroup.value = group;
  await fetchShortLinksForGroup(group.gid);
};

// 显示模态框
const showModal = () => {
  modalVisible.value = true;
  groupName.value = ''; // 清空输入框
};

// 处理创建分组
const handleCreateGroup = async () => {
  if (!groupName.value.trim()) {
    message.warning('请输入分组名称');
    return;
  }

  try {
    confirmLoading.value = true;
    const username = localStorage.getItem('username');
    if (!username) {
      message.error('用户未登录');
      return;
    }

    const response = await createGroup({
      name: groupName.value.trim(),
      username: username
    });

    if (response.data.success) {
      message.success('创建分组成功');
      modalVisible.value = false;
      // 重新获取分组列表
      await initializeGroups();
    } else {
      message.error(response.data.message || '创建分组失败');
    }
  } catch (error) {
    console.error("创建分组失败:", error);
    message.error('创建分组失败');
  } finally {
    confirmLoading.value = false;
  }
};

// 开始编辑
const startEdit = (group: GroupData) => {
  editingGroup.value = group;
  editingName.value = group.name;
  // 等待DOM更新后聚焦输入框
  setTimeout(() => {
    editInput.value?.focus();
  }, 0);
};

// 确认编辑
const confirmEdit = async () => {
  if (!editingGroup.value || !editingName.value.trim()) {
    message.warning('分组名称不能为空');
    return;
  }

  try {
    const username = localStorage.getItem('username');
    if (!username) {
      message.error('用户未登录');
      return;
    }

    const response = await updateGroup({
      gid: editingGroup.value.gid,
      name: editingName.value.trim(),
      username: username
    });

    if (response.data.success) {
      message.success('更新分组成功');
      // 重新获取分组列表
      await initializeGroups();
    } else {
      message.error(response.data.message || '更新分组失败');
    }
  } catch (error) {
    console.error("更新分组失败:", error);
    message.error('更新分组失败');
  } finally {
    editingGroup.value = null;
  }
};

// 取消编辑
const cancelEdit = () => {
  editingGroup.value = null;
  editingName.value = '';
};

// 开始拖拽
const handleDragStart = (event: DragEvent, group: GroupData) => {
  isDragging.value = true;
  draggingGroup.value = group;
  // 设置拖拽时的透明度
  if (event.target instanceof HTMLElement) {
    event.target.style.opacity = '0.5';
  }
};

// 结束拖拽
const handleDragEnd = (event: DragEvent) => {
  isDragging.value = false;
  draggingGroup.value = null;
  // 恢复透明度
  if (event.target instanceof HTMLElement) {
    event.target.style.opacity = '1';
  }
};

// 添加 dragover 处理函数
const handleDragOver = (event: DragEvent) => {
  event.preventDefault();
  event.dataTransfer!.dropEffect = 'move'; // 将默认的 copy 效果改为 move
};

// 处理删除
const handleDrop = async (event: DragEvent) => {
  event.preventDefault();
  if (!draggingGroup.value) return;

  try {
    const username = localStorage.getItem('username');
    if (!username) {
      message.error('用户未登录');
      return;
    }

    const response = await deleteGroup({
      gid: draggingGroup.value.gid,
      username: username
    });

    if (response.data.success) {
      message.success('删除分组成功');
      // 重新获取分组列表
      await initializeGroups();
    } else {
      message.error(response.data.message || '删除分组失败');
    }
  } catch (error) {
    console.error("删除分组失败:", error);
    message.error('删除分组失败');
  } finally {
    isDragging.value = false;
    draggingGroup.value = null;
  }
};

// 修改原有的 onMounted
onMounted(() => {
  initializeGroups();
});
</script>

<style>
.menu_content {
  display: flex;
  align-items: center; /* 让子元素在垂直方向上居中对齐 */
  margin-top: 15px;
}

.title {
  margin-left: 20px;
  font-size: 18px; /* 根据需要调整字体大小 */
  line-height: 1; /* 让文本与按钮保持一致高度 */
  display: flex;
  align-items: center; /* 确保文字垂直居中 */
}

.createGroup {
  margin-left: auto;
  font-size: 24px; /* 让按钮里的内容更大 */
  padding: 10px 20px; /* 调整按钮的整体大小 */
  display: flex;
  align-items: center; /* 垂直居中 */
  justify-content: center; /* 水平居中 */
}

.groupMenu{
  margin-top: 10px;
}

.group-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
}

.edit-actions {
  display: flex;
  gap: 8px;
}

.edit-icon {
  cursor: pointer;
  padding: 4px;
}

.edit-icon:hover {
  color: #1890ff;
}

/* 确认图标样式 */
.anticon-check {
  color: #52c41a;
}

/* 取消图标样式 */
.anticon-close {
  color: #ff4d4f;
}

.aside-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  position: relative;
}

.trash-zone {
  position: absolute;
  bottom: 20px;
  left: 20px;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s;
  cursor: pointer;
}

.trash-zone .anticon {
  font-size: 24px;  /* 稍微放大一点图标 */
  color: #1890ff;  /* 默认使用 antd 的主题蓝色 */
  transition: all 0.3s;
}

.trash-zone:hover .anticon,
.trash-active .anticon {
  color: #ff4d4f;  /* 悬停和拖动时变为红色 */
  transform: scale(1.2);  /* 稍微放大效果 */
}

/* 拖拽时的样式 */
.group-item {
  cursor: move;
  user-select: none;
}

[draggable="true"] {
  cursor: move;
}
</style>
