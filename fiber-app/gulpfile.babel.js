import gulp from "gulp";
import del from "del";
import image from "gulp-image";
import bro from "gulp-bro";
import babelify from "babelify";

const routes = {
  scss: {
    watch: "assets/scss/*",
    src: "assets/scss/*.scss",
    dest: "static/css",
  },
  img: {
    src: "assets/img/*",
    dest: "static/img",
  },
  js: {
    watch: "assets/js/**/*.js",
    src: "assets/js/main.js",
    dest: "static/js",
  },
};

// clean all object
const clean = () => del(["static/", ".publish"]);

// build scss
const scss = () => {
  const postCSS = require("gulp-postcss");
  const sass = require("gulp-sass")(require("sass"));
  const minify = require("gulp-csso");
  return gulp
    .src(routes.scss.src)
    .pipe(sass().on("error", sass.logError))
    .pipe(postCSS([require("tailwindcss"), require("autoprefixer")]))
    .pipe(minify())
    .pipe(gulp.dest(routes.scss.dest));
};

// zip image
const img = () =>
  gulp.src(routes.img.src).pipe(image()).pipe(gulp.dest(routes.img.dest));

// build JS
const js = () =>
  gulp.src(routes.js.src).pipe(
    bro({
      transform: [
        babelify.configure({ presets: ["@babel/preset-env"] }),
        ["uglifyify", { global: true }],
      ],
    })
  );

const watch = () => {
  gulp.watch(routes.img.src, img);
  gulp.watch(routes.scss.watch, scss);
  gulp.watch(routes.js.watch, js);
};

const prepare = gulp.series([clean, img]);
const assets = gulp.series([scss, js]);
const live = gulp.series([watch]);
export const build = gulp.series([prepare, assets]);
export const dev = gulp.series([build, live]);
