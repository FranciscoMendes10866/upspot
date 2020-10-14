<template>
  <div class="d-flex flex-wrap">
    <div
      v-for="event in events"
      :key="event.id"
      class="card border-0 shadow w mb-4 ml-4"
    >
      <span class="badge bg-white text-dark shadow bdg">{{ event.type }}.</span>
      <img :src="event.img" class="card-img-top" alt="..." />
      <div class="card-body">
        <h5 class="card-title">{{ event.title }}</h5>
        <small class="card-title">{{ event.date }}</small>
        <div class="mt-4 d-flex justify-content-between align-items-center">
          <nuxt-link :to="`/${event.id}`" class="btn btn-outline-primary"
            >View event</nuxt-link
          >
          <button
            v-if="event.attends_number != event.max"
            class="btn btn-light"
            disabled
          >
            {{ event.attends_number }} /
            <span class="font-weight-bold">{{ event.max }}</span>
          </button>
          <button v-else class="btn btn-light font-weight-bold" disabled>
            Full
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  // eslint-disable-next-line vue/require-prop-types
  props: ['events'],
})
</script>

<style scoped>
.w {
  width: 18rem;
}
.bdg {
  position: absolute;
  padding: 5px;
  top: 0;
  margin-top: 10px;
  margin-left: 10px;
}
</style>
