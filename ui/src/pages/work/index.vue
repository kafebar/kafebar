<script setup lang="ts">
const input = ref('')
const profession = ref(0)

const activate = computed(() => {
  if (input.value === '' || profession.value === 0)
    return false
  else
    return true
})

function setProfession(prof: number) {
  profession.value = prof
}
const kelnerActive = computed(() => {
  if (profession.value === 1)
    return 'bg-secondary border-primary'
  else
    return ''
})
const baristaActive = computed(() => {
  if (profession.value === 2)
    return 'bg-secondary'
  else
    return ''
})
function redirect() {
  if (activate.value)
    return 0
}
const egg = ref(0)
function iterate() {
  egg.value++
}
const easter = computed(() => {
  return egg.value < 10
})
</script>

<template>
  <div h-full w-full px-8>
    <div h-full flex flex-col items-center>
      <div mt-16 pb-10>
        <img v-show="easter" h-100 rounded-lg src="../../../public/coffee-gif.webp" @click="iterate()">
        <img v-show="!easter" h-100 rounded-lg src="../../../public/cat.gif">
      </div>
      <div h-full w-full flex flex-col>
        <div md:m-auto>
          <div my-5 md:pb-5>
            <p py-2 text-xl>
              Podaj PIN:
            </p>
            <input v-model="input" type="password" h-12 w-full border-b-2 border-l-2 border-secondary rounded-lg bg-primary px-4 py-3 md:w-500px>
          </div>
          <p py-2 text-xl>
            Logujesz się jako:
          </p>
          <div w-full flex flex-row border-b-2 border-l-2 border-secondary rounded-lg bg-primary md:mb-20 md:w-500px>
            <div h-25 w-full flex items-center justify-evenly p-4 @click="setProfession(1)">
              <div flex flex-row items-center>
                <div mr-4 h-5 w-5 border-2 border-secondary rounded-full :class="kelnerActive" />
                <p text-xl font-bold>
                  Kasjer
                </p>
              </div>
            </div>
            <div h-25 w-full flex items-center justify-evenly p-4 @click="setProfession(2)">
              <div flex flex-row items-center>
                <div mr-4 h-5 w-5 border-2 border-secondary rounded-full :class="baristaActive" />
                <p text-xl font-bold>
                  Barista
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
      <DefaultButton :active="activate" text="Zaloguj" fixed bottom-20 left-3 max-w-300px w-full @click="redirect()" />
    </div>
  </div>
</template>
