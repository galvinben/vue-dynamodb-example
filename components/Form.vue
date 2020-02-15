<template>
  <v-card class="card">
    <v-list dense>
      <v-subheader>Things:</v-subheader>
      <v-list-item-group color="primary">
        <v-list-item v-for="(item, i) in list" :key="i">
          <v-list-item-content>
            <v-list-item-title v-html="item.Name['S']"></v-list-item-title>
            <v-list-item-subtitle
              v-html="item.Info['S']"
            ></v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  </v-card>
</template>

<script>
export default {
  props: ['list']
}
</script>

<style scoped>
.card {
  width: 300px;
}
</style>
