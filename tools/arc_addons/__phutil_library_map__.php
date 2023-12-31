<?php

/**
 * This file is automatically generated. Use 'arc liberate' to rebuild it.
 *
 * @generated
 * @phutil-library-version 2
 */
phutil_register_library_map(array(
  '__library_version__' => 2,
  'class' => array(
    'ArcanistClangFormatLinter' => 'clang_format/lint/ArcanistClangFormatLinter.php',
    'ArcanistESLintLinter' => 'js/lint/ArcanistESLintLinter.php',
    'ArcanistGoImportsLinter' => 'pixielabs/lint/ArcanistGoImportsLinter.php',
    'ArcanistGolangCiLinter' => 'pixielabs/lint/ArcanistGolangCiLinter.php',
    'ArcanistPixieTextLinter' => 'pixielabs/lint/ArcanistPixieTextLinter.php',
    'ArcanistProtoBreakCheckLinter' => 'pixielabs/lint/ArcanistProtoBreakCheckLinter.php',
    'ArcanistShellCheckLinter' => 'shellcheck/lint/ArcanistShellCheckLinter.php',
    'ArcanistShellCheckLinterTestCase' => 'shellcheck/lint/__tests__/ArcanistShellCheckLinterTestCase.php',
  ),
  'function' => array(),
  'xmap' => array(
    'ArcanistClangFormatLinter' => 'ArcanistExternalLinter',
    'ArcanistESLintLinter' => 'ArcanistExternalLinter',
    'ArcanistGoImportsLinter' => 'ArcanistExternalLinter',
    'ArcanistGolangCiLinter' => 'ArcanistExternalLinter',
    'ArcanistPixieTextLinter' => 'ArcanistLinter',
    'ArcanistProtoBreakCheckLinter' => 'ArcanistExternalLinter',
    'ArcanistShellCheckLinter' => 'ArcanistExternalLinter',
    'ArcanistShellCheckLinterTestCase' => 'PhutilTestCase',
  ),
));
