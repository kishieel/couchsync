function(doc, req) {
    doc = doc || {};
    doc.updatedAt = new Date().toISOString();

    return [doc, 'updated'];
}