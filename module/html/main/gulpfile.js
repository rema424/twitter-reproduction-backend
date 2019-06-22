// Gulp packages
const gulp = require("gulp"),
  util = require("gulp-util"),
  del = require("del"),
  livereload = require("gulp-livereload"),
  extReplace = require("gulp-ext-replace"),
  htmlmin = require("gulp-htmlmin"),
  less = require("gulp-less"),
  postcss = require("gulp-postcss"),
  autoprefixer = require("autoprefixer"),
  autorem = require("autorem"),
  cssnano = require("cssnano"),
  uglify = require("gulp-uglify"),
  concat = require("gulp-concat"),
  imagemin = require("gulp-imagemin"),
  jpegRecompress = require("imagemin-jpeg-recompress"),
  pngQuant = require("imagemin-pngquant"),
  gifsicle = require("imagemin-gifsicle"),
  svgo = require("imagemin-svgo"),
  webp = require("imagemin-webp");

// HTML
function minifyHTML() {
  const src = "static/src/**/*.html",
    dest = "static/dist";

  return gulp
    .src(src)
    .pipe(
      htmlmin({
        collapseWhitespace: true,
        ignoreCustomFragments: [/{{[\s\S\.]*?}}/, /{%[\s\S\.]*?%}/],
        minifyCSS: true,
        minifyJS: true,
        // preserveLineBreaks: false,
        removeAttributeQuotes: true,
        removeComments: true,
        removeEmptyElements: true,
        removeOptionalTags: true,
        removeRedundantAttributes: true,
        removeScriptTypeAttributes: true,
        removeStyleLinkTypeAttributes: true,
        sortAttributes: true,
        sortClassName: true,
        trimCustomFragments: true,
        useShortDoctype: true
      })
    )
    .pipe(gulp.dest(dest))
    .pipe(livereload());
}

gulp.task(minifyHTML);

// CSS
function buildCSS() {
  const src = "static/src/less/main.less",
    dest = "static/dist/css";

  return gulp
    .src(src)
    .pipe(
      less().on("error", function(err) {
        util.log(err);
        this.emit("end");
      })
    )
    .pipe(
      postcss([
        autoprefixer({
          browsers: ["last 4 versions"]
        }),
        autorem(),
        cssnano()
      ])
    )
    .pipe(gulp.dest(dest))
    .pipe(livereload());
}

gulp.task(buildCSS);

// JavaScript
function uglifyJS() {
  const src = "static/src/js/**/*.js",
    dest = "static/dist/js";

  return gulp
    .src(src)
    .pipe(uglify())
    .pipe(gulp.dest(dest))
    .pipe(livereload());
}

gulp.task(uglifyJS);

function concatJS() {
  const src = ["static/dist/**/*.js", "!dist/js/scripts.js"],
    dest = "static/dist/js",
    concatScript = "scripts.js";

  return gulp
    .src(src)
    .pipe(concat(concatScript))
    .pipe(gulp.dest(dest))
    .pipe(livereload());
}

gulp.task(concatJS);

// Images
function imageminMain() {
  const src = "static/src/img/**/*.{png,jpg,svg,gif}",
    dest = "static/dist/img";

  return gulp
    .src(src, {
      since: gulp.lastRun("imageminMain")
    })
    .pipe(
      imagemin([
        jpegRecompress({
          max: 90
        }),
        pngQuant({
          quality: "45-90"
        }),
        gifsicle(),
        svgo()
      ])
    )
    .pipe(gulp.dest(dest))
    .pipe(livereload());
}

gulp.task(imageminMain);

// imagemin - WebP
function imageminWebP() {
  const src = "static/src/img/**/*.{jpg,png}",
    dest = "static/dist/img";

  return gulp
    .src(src, {
      since: gulp.lastRun("imageminWebP")
    })
    .pipe(
      imagemin([
        webp({
          quality: 65
        })
      ])
    )
    .pipe(extReplace(".webp"))
    .pipe(gulp.dest(dest))
    .pipe(livereload());
}

gulp.task(imageminWebP);

// Watch
function watch() {
  livereload.listen();

  gulp.watch("static/src/**/*.html", minifyHTML);
  gulp.watch("static/src/less/**/*.less", buildCSS);
  gulp.watch("static/src/js/**/*.js", gulp.series(uglifyJS, concatJS));
  gulp.watch("static/src/img/**/*.{png,jpg,svg,gif}", gulp.parallel(imageminMain, imageminWebP));
}

gulp.task("default", watch);

// Clean
function clean() {
  return del(["static/dist"]);
}

gulp.task(clean);

gulp.task(
  "build",
  gulp.parallel(minifyHTML, buildCSS, uglifyJS, imageminMain, imageminWebP, gulp.series(uglifyJS, concatJS))
);
