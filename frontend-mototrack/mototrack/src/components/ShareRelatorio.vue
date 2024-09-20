<template>
  <section id="shareRelatorio" class="content-section">
    <h2>COMPARTILHAR RELATÓRIO</h2>
    <hr>
    <!-- Content for sharing reports -->
    <button @click="generateRelatorio">Gerar Relatório</button>
    <div v-if="relatorio">
      <h3>Relatório</h3>
      <pre>{{ relatorio }}</pre>
    </div>
  </section>
</template>

<script>
export default {
  props: ['userId'], // Recebe o userId como prop
  data() {
    return {
      relatorio: null
    };
  },
  methods: {
    generateRelatorio() {
      const manutencoes = JSON.parse(localStorage.getItem('manutencoes')) || [];
      const relatorio = manutencoes.filter(m => m.userId === this.userId); // Filtra as manutenções pelo userId
      this.relatorio = JSON.stringify(relatorio, null, 2);
    }
  }
};
</script>

<style scoped>
.content-section {
  padding: 20px;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

button {
  padding: 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}

pre {
  text-align: left;
  background-color: #fff;
  padding: 10px;
  border-radius: 5px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}
</style>
