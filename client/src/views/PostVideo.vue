<template>
  <div class="post-video">
    <h2>欢迎投稿：</h2>
    <el-form ref="form" :model="form" label-width="80px">
      <el-form-item label="标题">
        <el-input v-model="form.title"></el-input>
      </el-form-item>

      <el-form-item label="上传视频">
        <el-upload
          class="upload-demo"
          action="/api/v1/upload/video"
          :before-upload="beforeVideoUpload"
          :on-success="handleVideoUploadSuccess"
          :show-file-list="true"
          :on-change="handleVideoUploadChange"
          :file-list="fileList">
          <el-button size="small" type="primary">点击上传</el-button>
          <div slot="tip" class="el-upload__tip">只能上传mp4文件，且不超过512M</div>
        </el-upload>
      </el-form-item>

      <el-form-item label="描述">
        <el-input type="textarea" v-model="form.info"></el-input>
      </el-form-item>

      <el-form-item label="视频封面">
        <el-upload
          class="avatar-uploader"
          label="描述"
          action="/api/v1/upload/image"
          ref="upload"
          :before-upload="beforeAvatarUpload"
          :on-success="handleAvatarUploadSuccess"
          :show-file-list="false">
          <img v-if="imageUrl" :src="imageUrl" class="avatar">
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
          <div class="el-upload__tip" slot="tip">只能上传png文件，且不超过500kb</div>
        </el-upload>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">立即创建</el-button>
      </el-form-item>

    </el-form>
  </div>
</template>

<script>
import * as API from '@/api/video/';
// import uplpadAPI from '@/api/upload/';

export default {
  name: 'PostVideo',
  data() {
    return {
      fileList: [],
      imageUrl: '',
      dialogImageUrl: '',
      dialogVisible: false,
      form: {
        title: '',
        info: '',
        url: '',
        avatar: '',
      },
    };
  },
  methods: {
    beforeVideoUpload(file) {
      const isVideo = (file.type === 'video/mp4');
      const isLt512M = file.size / 1024 / 1024 < 512;
      if (!isVideo) {
        this.$message.error('上传视频是mp4!');
        this.fileList = [];
      }
      if (!isLt512M) {
        this.$message.error('上传头像图片大小不能超过 512MB!');
        this.fileList = [];
      }
      return isVideo && isLt512M;
    },
    handleVideoUploadChange(file, nowFileList) {
      if (nowFileList.length > 0) {
        this.fileList = [nowFileList[nowFileList.length - 1]];
      }
    },
    handleVideoUploadSuccess(res) {
      if (res.status > 0) {
        this.$notify.error({
          title: '上传失败',
          message: res.msg,
        });
      } else {
        this.$notify({
          title: '上传成功！',
          message: '',
          type: 'success',
        });
        this.form.url = res.data.file_url;
      }
    },
    beforeAvatarUpload(file) {
      const isImage = (file.type === 'image/png' || file.type === 'image/jpeg');
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isImage) {
        this.$message.error('上传头像图片只能是图片!');
      }
      if (!isLt2M) {
        this.$message.error('上传头像图片大小不能超过 2MB!');
      }
      return isImage && isLt2M;
    },
    handleAvatarUploadSuccess(res) {
      if (res.status > 0) {
        this.$notify.error({
          title: '上传失败',
          message: res.msg,
        });
      } else {
        this.$notify({
          title: '上传成功！',
          message: '',
          type: 'success',
        });
        this.imageUrl = res.data.file_url;
        this.form.avatar = res.data.file_url;
      }
    },
    onSubmit() {
      if (this.form.url === '') {
        this.$notify.error({
          title: '投稿失败',
          message: '请上传视频',
        });
        return false;
      }
      if (this.form.avatar === '') {
        this.$notify.error({
          title: '投稿失败',
          message: '请上传图片',
        });
        return false;
      }
      API.postVideo(this.form).then((res) => {
        if (res.status > 0) {
          this.$notify.error({
            title: '投稿失败',
            message: res.msg,
          });
        } else {
          this.$notify({
            title: '投稿成功',
            message: `您投稿的ID为${res.data.id}`,
            type: 'success',
          });
          this.$router.push({ name: 'home' });
        }
      }).catch((error) => {
        this.$notify.error({
          title: '网路错误，或者服务器宕机',
          message: error,
        });
      });
      return true;
    },
  },
  components: {
  },
};
</script>

<style>
  .avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    line-height: 178px;
    text-align: center;
  }
  .avatar {
    max-width: 178px;
    max-height: 178px;
    display: block;
  }
</style>
