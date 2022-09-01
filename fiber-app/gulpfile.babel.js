import gulp from "gulp";
import bro from "gulp-bro";
import babelify from "babelify";
import tailwindCSS from "tailwindcss";
import autoprefixer from "autoprefixer";
import postCSS from "gulp-postcss";
import minify from "gulp-csso";
import del from "del";

const sass = require("gulp-sass")(require("sass"));

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
const clean = () => del(["static/"]);

// build scss
const scss = () => {
  return gulp
    .src(routes.scss.src)
    .pipe(sass().on("error", sass.logError))
    .pipe(postCSS([tailwindCSS, autoprefixer]))
    .pipe(minify())
    .pipe(gulp.dest(routes.scss.dest));
};

// zip image
// const img = () =>
//   gulp.src(routes.img.src).pipe(image()).pipe(gulp.dest(routes.img.dest));

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
  // gulp.watch(routes.img.src, img);
  gulp.watch(routes.scss.watch, scss);
  gulp.watch(routes.js.watch, js);
};

const prepare = gulp.series([clean]);
const assets = gulp.series([scss, js]);
const live = gulp.series([watch]);
export const build = gulp.series([prepare, assets]);
export const dev = gulp.series([build, live]);
