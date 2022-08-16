<template>
  <div>
    <p v-if="!isDefault">
      <label for="groupname">{{ $t("settings.groupname") }}</label>
      <input
        class="input input--block"
        type="text"
        v-model="group.groupname"
        id="groupname"
      />
    </p>

