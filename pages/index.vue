<template>
  <Form :list="list"></Form>
</template>

<script>
import Form from '~/components/Form'

export default {
  components: {
    Form
  },

  data: () => ({
    list: []
  }),

  mounted() {
    this.$axios
      .$get('https://6n7w1r95v5.execute-api.eu-west-2.amazonaws.com/dev/list')
      .then(response => {
        console.log(response)
        this.list = JSON.parse(response.body)
      })
  }
}
</script>
