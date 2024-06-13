<template>
  <div class="container">
    <h1>Pack Calculator</h1>
    <form @submit.prevent="calculatePacks">
      <label for="items">Number of Items:</label>
      <input type="number" v-model="items" min="1" required>
      <button type="submit">Calculate Packs</button>
    </form>
    <div v-if="packs.length">
      <h2>Required Packs:</h2>
      <ul>
        <li v-for="pack in packs" :key="pack.size">{{ pack.size }} items</li>
      </ul>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      items: null,
      packs: []
    };
  },
  methods: {
    async calculatePacks() {
      try {
        const response = await axios.get(`/api/packs?items=${this.items}`);
        this.packs = response.data;
      } catch (error) {
        console.error('Error fetching packs:', error);
      }
    }
  }
};
</script>

<style scoped>
.container {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  text-align: center;
}

h1 {
  margin-bottom: 20px;
}

form {
  margin-bottom: 20px;
}

input[type="number"] {
  padding: 8px;
  margin-right: 10px;
  width: 200px;
}

button {
  padding: 8px 16px;
  background-color: #007BFF;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  padding: 8px 0;
}
</style>