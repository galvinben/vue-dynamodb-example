<template>
  <div>
    <v-form ref="form" lazy-validation>
      <v-text-field
        v-model="name"
        :counter="20"
        :rules="nameRules"
        label="Name"
        required
      ></v-text-field>

      <v-text-field
        v-model="info"
        :rules="infoRules"
        label="Info"
        required
      ></v-text-field>

      <v-btn :disabled="!valid" color="success" class="mr-4" @click="add">
        Add
      </v-btn>
    </v-form>
    <p v-if="success">Success</p>
    <p v-if="failed">Failed</p>
  </div>
</template>

<script>
export default {
  data: () => ({
    valid: true,
    name: '',
    nameRules: [
      v => !!v || 'Name is required',
      v => (v && v.length <= 20) || 'Name must be less than 20 characters'
    ],
    info: '',
    infoRules: [v => !!v || 'Info is required'],
    success: false,
    failed: false
  }),
  methods: {
    add() {
      if (this.$refs.form.validate()) {
        this.$axios
          .$post(
            'https://6n7w1r95v5.execute-api.eu-west-2.amazonaws.com/dev/list',
            {
              name: this.name,
              info: this.info,
              category: 'category'
            }
          )
          .then(response => {
            console.log(response)
            this.success = true
          })
          .catch(error => {
            console.log(error.response)
            this.failed = true
          })
      }
    }
  }
}
</script>
