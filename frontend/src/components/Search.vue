<template>
  <div>
    <input @keyup="searchYachts" v-model.trim="query" type="text" class="form-control" placeholder="Search...">
    <div class="mt-4">
      <Yacht v-for="yacht in yachts" :key="yacht.id" :yacht="yacht" />
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex';
import Yacht from '@/components/Yacht';

export default {
  data() {
    return {
      query: '',
    };
  },
  computed: mapState({
    yachts: (state) => state.searchResults,
  }),
  methods: {
    searchYachts() {
      if (this.query != this.lastQuery) {
        this.$store.dispatch('searchYachts', this.query);
        this.lastQuery = this.query;
      }
    },
  },
  components: {
    Yacht,
  },
};
</script>
