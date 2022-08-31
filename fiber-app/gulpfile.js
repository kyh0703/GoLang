const gulp = require("gulp");

const routes = {
  css: {
    watch: "assets/scss/*",
    src: "assets/scss/*.scss",
    dest: "static/css",
  },
};

const css = () => {
  const postCSS = require("gulp-postcss");
  const sass = require("gulp-sass")(require("sass"));
  const minify = require("gulp-csso");
  return gulp
    .src(routes.css.src)
    .pipe(sass().on("error", sass.logError))
    .pipe(postCSS([require("tailwindcss"), require("autoprefixer")]))
    .pipe(minify())
    .pipe(gulp.dest(routes.css.dest));
};

exports.default = css;
