<template>
  <div>
    <List :list="list"></List>
    <Add></Add>
  </div>
</template>

<script>
import List from '~/components/List'
import Add from '~/components/Add'

export default {
  components: {
    List,
    Add
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
