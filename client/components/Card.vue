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
          <nuxt-link to="/" class="btn btn-outline-primary"
            >View event</nuxt-link
          >
          <button
            v-if="event.current != event.max"
            class="btn btn-light"
            disabled
          >
            {{ event.current }} /
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
  data: () => ({
    events: [
      {
        id: 1,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Online event',
        max: 35,
        current: 12,
      },
      {
        id: 2,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Online event',
        max: 50,
        current: 4,
      },
      {
        id: 3,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Online event',
        max: 12,
        current: 12,
      },
      {
        id: 4,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Online event',
        max: 40,
        current: 27,
      },
      {
        id: 5,
        img:
          'https://secure.meetupstatic.com/photos/event/4/3/e/2/highres_492797378.jpeg',
        title: 'Agile Connect Day',
        date: 'Saturday, October 10, 2020',
        type: 'Online event',
        max: 200,
        current: 44,
      },
    ],
  }),
})
</script>
