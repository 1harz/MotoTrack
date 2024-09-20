
<template>
  <div class="fundo_control">
    <div id="app">
      <NavBar @changeSection="currentSection = $event" @logout="logout" class="sidebar" />
      <main class="content container equal-height">
        <WelcomeSection v-if="currentSection === 'welcomeSection'" />
        <RegisterMoto v-if="currentSection === 'registroMoto'" :userId="userId" />
        <ListMotos v-if="currentSection === 'listMotos'" :userId="userId" />
        <RegisterManutencao v-if="currentSection === 'registerManutencao'" :userId="userId" />
        <ConsultManutencao v-if="currentSection === 'consultManutencao'" :userId="userId" />
        <ShareRelatorio v-if="currentSection === 'shareRelatorio'" :userId="userId" />
      </main>
    </div>
  </div>
  <div class="account-container" @click="redirectUser">
    <div class="account">
      <span>Minha Conta</span>
      <img src="@/assets/user_photo.png" alt="Foto" class="account-photo">
    </div>
  </div>
</template>

<script>
import NavBar from '../components/NavBar.vue';
import WelcomeSection from '../components/WelcomeSection.vue';
import RegisterMoto from '../components/RegisterMoto.vue';
import ListMotos from '../components/ListMotos.vue';
import RegisterManutencao from '../components/RegisterManutencao.vue';
import ConsultManutencao from '../components/ConsultManutencao.vue';
import ShareRelatorio from '../components/ShareRelatorio.vue';

export default {
  components: {
    NavBar,
    WelcomeSection,
    RegisterMoto,
    ListMotos,
    RegisterManutencao,
    ConsultManutencao,
    ShareRelatorio
  },
  data() {
    return {
      currentSection: 'welcomeSection',
      userId: localStorage.getItem('userId') // Obtém o ID do usuário do LocalStorage
    };
  },
  methods: {
    redirectUser() {
      this.$router.push('/conta');
    },
    logout() {
      localStorage.removeItem('userId'); // Remove o ID do usuário do LocalStorage
      this.$router.push('/login'); // Redireciona para a página de login
    }
  }
};
</script>

<style scoped>
.fundo_control {
  background: url('@/assets/motogp.jpg') no-repeat center center fixed;
  background-size: cover;
  height: 100vh;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

#app {
  display: flex;
  height: 100%;
  width: 100%;
  flex-direction: row;
  position: absolute;
  top: 9%;
  left: 10%;
}

.sidebar {
  flex: 0 0 250px;
  height: 80%;
}

.container {
  flex: 1;
  max-width: 1200px;
  background-color: #fff;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.8);
  border-radius: 10px;
  padding: 20px;
  display: flex;
  flex-direction: column;
  overflow: auto;
  max-height: 90vh; /* Ajuste a altura máxima */
}

.account-container {
  position: absolute;
  top: 5px;
  right: 5px;
  display: flex;
  align-items: center;
  background-color: white;
  padding: 2px;
  border-radius: 10px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
}

.account {
  display: flex;
  align-items: center;
}

.account span {
  margin-right: 10px;
  font-size: 0.8em;
  font-weight: bold;
  color: #0044cc;
}

.account-photo {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  border: 2px solid #0044cc;
}

.equal-height {
  min-height: 40%;
  max-height: 80vh;
  overflow: auto;
}
</style>