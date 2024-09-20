<template>
    <div>
      <RegisterMoto :userId="userId"/>
    </div>
  </template>
  
  <script>
  import RegisterMoto from './RegisterMoto.vue';
  
  export default {
    components: {
      RegisterMoto
    },
    data() {
      return {
        userId: localStorage.getItem('userId') // Obter o userId do localStorage
      };
    }
  };
  </script>
  