<template>
  <section id="consultarManutencao" class="content-section">
    <h2>LISTA DE MANUTENÇÕES</h2>
    <hr>
    <button @click="clearManutencaoList">Limpar Lista de Manutenções</button> <br> <br>
    <div v-if="manutencoes.length === 0">
      <p>Nenhuma manutenção disponível.</p>
    </div>
    <div v-for="(manutencao, index) in manutencoes" :key="index" class="manutencao-card">
      <h3>{{ manutencao.tm }}</h3>
      <p><strong>Data:</strong> {{ manutencao.date }}</p>
      <p><strong>Valor:</strong> {{ manutencao.value }}</p>
      <p><strong>Descrição:</strong> {{ manutencao.description }}</p>
    </div>
  </section>
</template>

<script>
import axios from 'axios';

export default {
  props: ['userId'], // Recebe o userId como prop
  data() {
    return {
      manutencoes: [] // Lista de manutenções
    };
  },
  created() {
    this.loadManutencaoList(); // Carrega a lista de manutenções ao criar o componente
  },
  methods: {
    // Função para carregar a lista de manutenções
    async loadManutencaoList() {
      try {
        const response = await axios.get(`http://localhost:8080/maintenances?userId=${this.userId}`);
        this.manutencoes = response.data;
      } catch (error) {
        console.error('Error loading maintenances:', error);
        alert('Erro ao carregar a lista de manutenções.');
      }
    },
    // Função para limpar a lista de manutenções
    clearManutencaoList() {
      // Enviar uma requisição DELETE para limpar as manutenções do usuário
      axios.delete(`http://localhost:8080/maintenances?userId=${this.userId}`)
        .then(() => {
          this.manutencoes = [];
        })
        .catch(error => {
          console.error('Error clearing maintenances:', error);
          alert('Erro ao limpar a lista de manutenções.');
        });
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

.manutencao-card {
  margin-bottom: 20px;
  padding: 15px;
  border: 1px solid #ddd;
  border-radius: 8px;
  background-color: #fff;
}

.manutencao-card h3 {
  margin: 0 0 10px 0;
}

button {
  padding: 10px;
  margin-top: 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
