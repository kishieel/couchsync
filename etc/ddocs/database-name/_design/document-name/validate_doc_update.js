function(newDoc, oldDoc, userCtx, secObj) {
    const whitelist = ["_admin", "_writer"];

    if (!whitelist.some(role => userCtx.roles.includes(role))) {
        throw ({forbidden: "You are not authorized to write to this database"});
    }
}
