<template>
  <section class="main">
    <div class="container main-container is-white left-main">
      <div class="left-container">
        <project-list :projects="projectPage.results" />
        <pagination :page="projectPage.page" url-prefix="/projects/" />
      </div>
      <div class="right-container">
        <site-notice />

        <div class="ad">
          <!-- 展示广告 -->
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  async asyncData({ $axios, params }) {
    const [projectPage] = await Promise.all([
      $axios.get('/api/project/projects', {
        params: {
          page: params.page || 1,
        },
      }),
    ])
    return {
      projectPage,
    }
  },
  head() {
    return {
      title: this.$siteTitle('开源项目'),
      meta: [
        {
          hid: 'description',
          name: 'description',
          content: this.$siteDescription(),
        },
        { hid: 'keywords', name: 'keywords', content: this.$siteKeywords() },
      ],
    }
  },
}
</script>

<style lang="scss" scoped></style>
