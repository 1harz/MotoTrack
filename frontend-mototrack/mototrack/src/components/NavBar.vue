<template>
  <nav class="menu">
    <div class="logo-container" @click="changeSection('welcomeSection')">
      <img src="@/assets/logo_mt.png" alt="Logo" class="logo">
    </div>
    <button @click="changeSection('registroMoto')">Registrar Moto</button>
    <button @click="changeSection('listMotos')">Lista de Motos</button>
    <button @click="changeSection('registerManutencao')">Registrar Manutenção</button>
    <button @click="changeSection('consultManutencao')">Lista de Manutenção</button>
    <!-- <button @click="changeSection('shareRelatorio')">Compartilhar Relatório</button> -->
    <button @click="logout">Logout</button> <!-- Adiciona botão de logout -->
  </nav>
</template>

<script>
export default {
  methods: {
    changeSection(section) {
      this.$emit('changeSection', section);
    },
    logout() {
      this.$emit('logout'); // Emite evento de logout
    }
  }
};
</script>

<style scoped>
.menu {
  flex: 0 0 15%;
  background-color: #1a1a1a;
  padding: 20px;
  color: white;
  display: flex;
  flex-direction: column;
  align-items: center;
  border-top-left-radius: 10px;
  border-bottom-left-radius: 10px;
  height: 70%;
}

.logo-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 20px;
  cursor: pointer;
  max-height: 150px;
  overflow: hidden;
}

.logo {
  max-width: 100%;
  height: auto;
  margin: 0 auto;
}

.menu button {
  background-color: #0044cc;
  color: white;
  border: none;
  padding: 10px 20px;
  margin: 10px 0;
  cursor: pointer;
  width: 100%;
  text-align: center;
  border-radius: 10px;
  transition: all 0.3s;
}

.menu button:hover {
  background-color: white;
  color: #0044cc;
  border: 2px solid #0044cc;
  transform: scale(1.05);
}
</style>