package models

const createUpdateTranslationsFunc = `
DROP FUNCTION IF EXISTS update_translations;
CREATE FUNCTION update_translations(ts translations[])
    RETURNS void
    LANGUAGE sql AS
$func$
    UPDATE translations
    SET content = s.content
    FROM unnest(ts) s 
    WHERE translations.uuid = s.uuid AND translations.locale = s.locale
$func$;
`
