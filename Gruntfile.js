module.exports = function(grunt) {
    'use strict';

    grunt.initConfig({
        subgrunt: {
            target0: {
                projects: {
                    'static': 'build'
                }
            }
        },
        changelog: {
            options: {
                repository: 'https://github.com/SII-Nantes/callForPaper',
                version: grunt.option("rev") != undefined ? grunt.option("rev") : "",
                from : grunt.option("from") 
            }
        }
    });

    grunt.loadNpmTasks('grunt-subgrunt');
    grunt.loadNpmTasks('grunt-conventional-changelog');
    grunt.registerTask('build', ['subgrunt']);
    grunt.registerTask('publish', ['changelog']);
}