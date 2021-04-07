<template>
  <div class="row py-3 justify-content-center" v-if="totalPages">
    <div class="col-auto">
      <nav aria-label="Page navigation">
        <ul class="pagination">
          <!--     現在、最新ページにいる場合、最新ページへ移動するリンク << を無効にする     -->
          <li class="page-item"
              :class="{'disabled': currentPageNum == newestPageNum}">
            <a class="page-link"
               href="#"
               v-on:click.prevent="setPage(newestPageNum);">
              <<
            </a>
          </li>

          <!--    現在、最新ページにいる場合、前ページへ移動するリンク < を無効にする      -->
          <li class="page-item"
              :class="{'disabled': currentPageNum == newestPageNum}">
            <a class="page-link" href="#"
               v-on:click.prevent="setPage(currentPageNum - 1);"
               :class="{'disable':currentPageNum == newestPageNum}">
              <
            </a>
          </li>

          <!--     ページセット     -->
          <li
            class="page-item"
            v-for="num in showPagesFix"
            :key="num"
            :class="{'active' : generatePageNum(num) == currentPageNum}"
          >
            <template v-if="generatePageNum(num) == currentPageNum">
              <span class="page-link">
                {{ generatePageNum(num) }}
              </span>
            </template>
            <a
              class="page-link"
              href="#"
              v-on:click.prevent="setPage(generatePageNum(num))"
              v-else
            >
              {{ generatePageNum(num) }}
            </a>
          </li>

          <!--     現在、最古ページにいる場合、次ページへ移動するリンク > を無効にする     -->
          <li class="page-item"
              :class="{'disabled': currentPageNum == totalPages}">
            <a class="page-link"
               href="#"
               v-on:click.prevent="setPage(currentPageNum + 1);">
              >
            </a>
          </li>

          <!--     現在、最古ページにいる場合、最古ページへ移動するリンク >> を無効にする     -->
          <li class="page-item"
              :class="{'disabled': currentPageNum == totalPages}">
            <a class="page-link"
               href="#"
               v-on:click.prevent="setPage(totalPages);">
              >>
            </a>
          </li>

        </ul>
      </nav>
    </div>
  </div>
</template>

<script>
export default {

  data() {
    return {
      currentPageNum: Number, //現在のページ
    };
  },

  props: {
    showPages: Number, //ページネーションを何件表示するか
    currentPage: Number, //現在のページ
    newestPageNum: 1, //最新ページ
    totalPages: Number, //総ページ数
  },

  mounted() {
    this.$set(this, "currentPageNum", this.currentPage);
  },

  computed: {
    //ページ番号を計算する
    generatePageNum() {
      return function (num) {
        let ajust = 1 + (this.showPages - 1) / 2;
        let result = num;

        //前ページがマイナスになる場合は1からはじめる
        if (this.currentPageNum > this.showPages / 2) {
          result = num + this.currentPageNum - ajust;
        }

        //後ページが最大ページを超える場合は最大ページを超えないようにする
        if (this.currentPageNum + this.showPages / 2 > this.totalPages) {
          result = this.totalPages - this.showPages + num;
        }

        //総ページ数が表示ページ数に満たない場合、連番そのまま
        if (this.totalPages <= this.showPages) {
          result = num;
        }

        return result;
      };
    },

    //総記事数が表示ページ数以下の場合に調整する
    showPagesFix() {
      return this.totalPages < this.showPages ? this.totalPages : this.showPages;
    },

  },

  methods: {
    //何ページ目を表示するか
    setPage(page) {
      //マイナスにならないようにする
      if (page <= 0) {
        this.$set(this, "currentPageNum", this.newestPageNum);
      }
      //最大ページを超えないようにする
      else if (page > this.totalPages) {
        this.$set(this, "currentPageNum", this.totalPages);
      } else {
        this.$set(this, "currentPageNum", page);
      }
      //親コンポーネントに現在のページを送る
      this.$emit("currentPage", this.currentPageNum);
    },
  },

  watch: {
    //ページネーションを複数設置したときの対応
    currentPage(val) {
      this.$set(this, "currentPageNum", this.currentPage);
    },
  },

};
</script>
