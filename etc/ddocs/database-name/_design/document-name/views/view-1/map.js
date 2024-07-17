function(doc) {
    if (doc.type === 'post') {
        emit(doc._id, doc);
    }
}