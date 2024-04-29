package generated

import (
	"time"
)

type Todo__CreateOneTodoInput struct {
	Title string `json:"title"`
}

type Todo__CreateOneTodoInternalInput struct {
	Title string `json:"title"`
}

type Todo__CreateOneTodoResponseData struct {
	Data Todo__CreateOneTodoResponseData_data `json:"data,omitempty"`
}

type Todo__CreateOneTodoResponseData_data struct {
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
}

type Todo__DeleteOneTodoInput struct {
	Id int64 `json:"id"`
}

type Todo__DeleteOneTodoInternalInput struct {
	Id int64 `json:"id"`
}

type Todo__DeleteOneTodoResponseData struct {
	Data Todo__DeleteOneTodoResponseData_data `json:"data,omitempty"`
}

type Todo__DeleteOneTodoResponseData_data struct {
	Id int64 `json:"id"`
}

type Todo__GetManyTodoInput struct {
}

type Todo__GetManyTodoInternalInput struct {
}

type Todo__GetManyTodoResponseData struct {
	Data []Todo__GetManyTodoResponseData_data `json:"data"`
}

type Todo__GetManyTodoResponseData_data struct {
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
}

type Todo__UpdateOneTodoInput struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type Todo__UpdateOneTodoInternalInput struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type Todo__UpdateOneTodoResponseData struct {
	Data Todo__UpdateOneTodoResponseData_data `json:"data,omitempty"`
}

type Todo__UpdateOneTodoResponseData_data struct {
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
}

type Todo__UpdateTodoCompletedInput struct {
	Completed bool  `json:"completed"`
	Id        int64 `json:"id"`
}

type Todo__UpdateTodoCompletedInternalInput struct {
	Completed bool  `json:"completed"`
	Id        int64 `json:"id"`
}

type Todo__UpdateTodoCompletedResponseData struct {
	Data Todo__UpdateTodoCompletedResponseData_data `json:"data,omitempty"`
}

type Todo__UpdateTodoCompletedResponseData_data struct {
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
}

type Blog_BoolFilter struct {
	Equals bool                   `json:"equals,omitempty"`
	Not    *Blog_NestedBoolFilter `json:"not,omitempty"`
}

type Blog_CategoryCreateNestedManyWithoutPostInput struct {
	Connect         *Blog_CategoryWhereUniqueInput                `json:"connect,omitempty"`
	ConnectOrCreate *Blog_CategoryCreateOrConnectWithoutPostInput `json:"connectOrCreate,omitempty"`
	Create          *Blog_CategoryCreateWithoutPostInput          `json:"create,omitempty"`
}

type Blog_CategoryCreateOrConnectWithoutPostInput struct {
	Create *Blog_CategoryCreateWithoutPostInput `json:"create"`
	Where  *Blog_CategoryWhereUniqueInput       `json:"where"`
}

type Blog_CategoryCreateWithoutPostInput struct {
	Name string `json:"name"`
}

type Blog_CategoryListRelationFilter struct {
	Every *Blog_CategoryWhereInput `json:"every,omitempty"`
	None  *Blog_CategoryWhereInput `json:"none,omitempty"`
	Some  *Blog_CategoryWhereInput `json:"some,omitempty"`
}

type Blog_CategoryWhereInput struct {
	AND  *Blog_CategoryWhereInput     `json:"AND,omitempty"`
	NOT  *Blog_CategoryWhereInput     `json:"NOT,omitempty"`
	OR   []*Blog_CategoryWhereInput   `json:"OR,omitempty"`
	Post *Blog_PostListRelationFilter `json:"Post,omitempty"`
	Id   *Blog_IntFilter              `json:"id,omitempty"`
	Name *Blog_StringFilter           `json:"name,omitempty"`
}

type Blog_CategoryWhereUniqueInput struct {
	AND  *Blog_CategoryWhereInput     `json:"AND,omitempty"`
	NOT  *Blog_CategoryWhereInput     `json:"NOT,omitempty"`
	OR   []*Blog_CategoryWhereInput   `json:"OR,omitempty"`
	Post *Blog_PostListRelationFilter `json:"Post,omitempty"`
	Id   int64                        `json:"id,omitempty"`
	Name string                       `json:"name,omitempty"`
}

type Blog_IntFilter struct {
	Equals int64                 `json:"equals,omitempty"`
	Gt     int64                 `json:"gt,omitempty"`
	Gte    int64                 `json:"gte,omitempty"`
	In     []int64               `json:"in,omitempty"`
	Lt     int64                 `json:"lt,omitempty"`
	Lte    int64                 `json:"lte,omitempty"`
	Not    *Blog_NestedIntFilter `json:"not,omitempty"`
	NotIn  []int64               `json:"notIn,omitempty"`
}

type Blog_IntNullableFilter struct {
	Equals int64                         `json:"equals,omitempty"`
	Gt     int64                         `json:"gt,omitempty"`
	Gte    int64                         `json:"gte,omitempty"`
	In     []int64                       `json:"in,omitempty"`
	Lt     int64                         `json:"lt,omitempty"`
	Lte    int64                         `json:"lte,omitempty"`
	Not    *Blog_NestedIntNullableFilter `json:"not,omitempty"`
	NotIn  []int64                       `json:"notIn,omitempty"`
}

type Blog_NestedBoolFilter struct {
	Equals bool                   `json:"equals,omitempty"`
	Not    *Blog_NestedBoolFilter `json:"not,omitempty"`
}

type Blog_NestedIntFilter struct {
	Equals int64                 `json:"equals,omitempty"`
	Gt     int64                 `json:"gt,omitempty"`
	Gte    int64                 `json:"gte,omitempty"`
	In     []int64               `json:"in,omitempty"`
	Lt     int64                 `json:"lt,omitempty"`
	Lte    int64                 `json:"lte,omitempty"`
	Not    *Blog_NestedIntFilter `json:"not,omitempty"`
	NotIn  []int64               `json:"notIn,omitempty"`
}

type Blog_NestedIntNullableFilter struct {
	Equals int64                         `json:"equals,omitempty"`
	Gt     int64                         `json:"gt,omitempty"`
	Gte    int64                         `json:"gte,omitempty"`
	In     []int64                       `json:"in,omitempty"`
	Lt     int64                         `json:"lt,omitempty"`
	Lte    int64                         `json:"lte,omitempty"`
	Not    *Blog_NestedIntNullableFilter `json:"not,omitempty"`
	NotIn  []int64                       `json:"notIn,omitempty"`
}

type Blog_NestedStringFilter struct {
	Contains   string                   `json:"contains,omitempty"`
	EndsWith   string                   `json:"endsWith,omitempty"`
	Equals     string                   `json:"equals,omitempty"`
	Gt         string                   `json:"gt,omitempty"`
	Gte        string                   `json:"gte,omitempty"`
	In         []string                 `json:"in,omitempty"`
	Lt         string                   `json:"lt,omitempty"`
	Lte        string                   `json:"lte,omitempty"`
	Not        *Blog_NestedStringFilter `json:"not,omitempty"`
	NotIn      []string                 `json:"notIn,omitempty"`
	StartsWith string                   `json:"startsWith,omitempty"`
}

type Blog_NestedStringNullableFilter struct {
	Contains   string                           `json:"contains,omitempty"`
	EndsWith   string                           `json:"endsWith,omitempty"`
	Equals     string                           `json:"equals,omitempty"`
	Gt         string                           `json:"gt,omitempty"`
	Gte        string                           `json:"gte,omitempty"`
	In         []string                         `json:"in,omitempty"`
	Lt         string                           `json:"lt,omitempty"`
	Lte        string                           `json:"lte,omitempty"`
	Not        *Blog_NestedStringNullableFilter `json:"not,omitempty"`
	NotIn      []string                         `json:"notIn,omitempty"`
	StartsWith string                           `json:"startsWith,omitempty"`
}

type Blog_PostCreateNestedManyWithoutUserInput struct {
	Connect         *Blog_PostWhereUniqueInput                `json:"connect,omitempty"`
	ConnectOrCreate *Blog_PostCreateOrConnectWithoutUserInput `json:"connectOrCreate,omitempty"`
	Create          *Blog_PostCreateWithoutUserInput          `json:"create,omitempty"`
}

type Blog_PostCreateOrConnectWithoutUserInput struct {
	Create *Blog_PostCreateWithoutUserInput `json:"create"`
	Where  *Blog_PostWhereUniqueInput       `json:"where"`
}

type Blog_PostCreateWithoutUserInput struct {
	Category  *Blog_CategoryCreateNestedManyWithoutPostInput `json:"Category,omitempty"`
	Content   string                                         `json:"content,omitempty"`
	Like      int64                                          `json:"like,omitempty"`
	Published bool                                           `json:"published"`
	Title     string                                         `json:"title"`
	View      int64                                          `json:"view,omitempty"`
}

type Blog_PostListRelationFilter struct {
	Every *Blog_PostWhereInput `json:"every,omitempty"`
	None  *Blog_PostWhereInput `json:"none,omitempty"`
	Some  *Blog_PostWhereInput `json:"some,omitempty"`
}

type Blog_PostWhereInput struct {
	AND       *Blog_PostWhereInput             `json:"AND,omitempty"`
	Category  *Blog_CategoryListRelationFilter `json:"Category,omitempty"`
	NOT       *Blog_PostWhereInput             `json:"NOT,omitempty"`
	OR        []*Blog_PostWhereInput           `json:"OR,omitempty"`
	User      *Blog_UserNullableRelationFilter `json:"User,omitempty"`
	AuthorId  *Blog_IntNullableFilter          `json:"authorId,omitempty"`
	Content   *Blog_StringNullableFilter       `json:"content,omitempty"`
	Id        *Blog_IntFilter                  `json:"id,omitempty"`
	Like      *Blog_IntNullableFilter          `json:"like,omitempty"`
	Published *Blog_BoolFilter                 `json:"published,omitempty"`
	Title     *Blog_StringFilter               `json:"title,omitempty"`
	View      *Blog_IntNullableFilter          `json:"view,omitempty"`
}

type Blog_PostWhereUniqueInput struct {
	AND       *Blog_PostWhereInput             `json:"AND,omitempty"`
	Category  *Blog_CategoryListRelationFilter `json:"Category,omitempty"`
	NOT       *Blog_PostWhereInput             `json:"NOT,omitempty"`
	OR        []*Blog_PostWhereInput           `json:"OR,omitempty"`
	User      *Blog_UserNullableRelationFilter `json:"User,omitempty"`
	AuthorId  *Blog_IntNullableFilter          `json:"authorId,omitempty"`
	Content   *Blog_StringNullableFilter       `json:"content,omitempty"`
	Id        int64                            `json:"id,omitempty"`
	Like      *Blog_IntNullableFilter          `json:"like,omitempty"`
	Published *Blog_BoolFilter                 `json:"published,omitempty"`
	Title     *Blog_StringFilter               `json:"title,omitempty"`
	View      *Blog_IntNullableFilter          `json:"view,omitempty"`
}

type Blog_ProfileCreateNestedOneWithoutUserInput struct {
	Connect         *Blog_ProfileWhereUniqueInput                `json:"connect,omitempty"`
	ConnectOrCreate *Blog_ProfileCreateOrConnectWithoutUserInput `json:"connectOrCreate,omitempty"`
	Create          *Blog_ProfileCreateWithoutUserInput          `json:"create,omitempty"`
}

type Blog_ProfileCreateOrConnectWithoutUserInput struct {
	Create *Blog_ProfileCreateWithoutUserInput `json:"create"`
	Where  *Blog_ProfileWhereUniqueInput       `json:"where"`
}

type Blog_ProfileCreateWithoutUserInput struct {
	Address string `json:"address"`
}

type Blog_ProfileNullableRelationFilter struct {
	Is    *Blog_ProfileWhereInput `json:"is,omitempty"`
	IsNot *Blog_ProfileWhereInput `json:"isNot,omitempty"`
}

type Blog_ProfileWhereInput struct {
	AND     *Blog_ProfileWhereInput   `json:"AND,omitempty"`
	NOT     *Blog_ProfileWhereInput   `json:"NOT,omitempty"`
	OR      []*Blog_ProfileWhereInput `json:"OR,omitempty"`
	User    *Blog_UserRelationFilter  `json:"User,omitempty"`
	Address *Blog_StringFilter        `json:"address,omitempty"`
	Id      *Blog_IntFilter           `json:"id,omitempty"`
	UserId  *Blog_IntFilter           `json:"userId,omitempty"`
}

type Blog_ProfileWhereUniqueInput struct {
	AND     *Blog_ProfileWhereInput   `json:"AND,omitempty"`
	NOT     *Blog_ProfileWhereInput   `json:"NOT,omitempty"`
	OR      []*Blog_ProfileWhereInput `json:"OR,omitempty"`
	User    *Blog_UserRelationFilter  `json:"User,omitempty"`
	Address *Blog_StringFilter        `json:"address,omitempty"`
	Id      int64                     `json:"id,omitempty"`
	UserId  int64                     `json:"userId,omitempty"`
}

type Blog_StringFilter struct {
	Contains   string                   `json:"contains,omitempty"`
	EndsWith   string                   `json:"endsWith,omitempty"`
	Equals     string                   `json:"equals,omitempty"`
	Gt         string                   `json:"gt,omitempty"`
	Gte        string                   `json:"gte,omitempty"`
	In         []string                 `json:"in,omitempty"`
	Lt         string                   `json:"lt,omitempty"`
	Lte        string                   `json:"lte,omitempty"`
	Not        *Blog_NestedStringFilter `json:"not,omitempty"`
	NotIn      []string                 `json:"notIn,omitempty"`
	StartsWith string                   `json:"startsWith,omitempty"`
}

type Blog_StringNullableFilter struct {
	Contains   string                           `json:"contains,omitempty"`
	EndsWith   string                           `json:"endsWith,omitempty"`
	Equals     string                           `json:"equals,omitempty"`
	Gt         string                           `json:"gt,omitempty"`
	Gte        string                           `json:"gte,omitempty"`
	In         []string                         `json:"in,omitempty"`
	Lt         string                           `json:"lt,omitempty"`
	Lte        string                           `json:"lte,omitempty"`
	Not        *Blog_NestedStringNullableFilter `json:"not,omitempty"`
	NotIn      []string                         `json:"notIn,omitempty"`
	StartsWith string                           `json:"startsWith,omitempty"`
}

type Blog_UserCreateInput struct {
	Post    *Blog_PostCreateNestedManyWithoutUserInput   `json:"Post,omitempty"`
	Profile *Blog_ProfileCreateNestedOneWithoutUserInput `json:"Profile,omitempty"`
	Age     int64                                        `json:"age,omitempty"`
	Country string                                       `json:"country,omitempty"`
	Email   string                                       `json:"email"`
}

type Blog_UserNullableRelationFilter struct {
	Is    *Blog_UserWhereInput `json:"is,omitempty"`
	IsNot *Blog_UserWhereInput `json:"isNot,omitempty"`
}

type Blog_UserRelationFilter struct {
	Is    *Blog_UserWhereInput `json:"is,omitempty"`
	IsNot *Blog_UserWhereInput `json:"isNot,omitempty"`
}

type Blog_UserWhereInput struct {
	AND     *Blog_UserWhereInput                `json:"AND,omitempty"`
	NOT     *Blog_UserWhereInput                `json:"NOT,omitempty"`
	OR      []*Blog_UserWhereInput              `json:"OR,omitempty"`
	Post    *Blog_PostListRelationFilter        `json:"Post,omitempty"`
	Profile *Blog_ProfileNullableRelationFilter `json:"Profile,omitempty"`
	Age     *Blog_IntNullableFilter             `json:"age,omitempty"`
	Country *Blog_StringNullableFilter          `json:"country,omitempty"`
	Email   *Blog_StringFilter                  `json:"email,omitempty"`
	Id      *Blog_IntFilter                     `json:"id,omitempty"`
}

type N01curd__create__createInput struct {
	Email string `json:"email,omitempty"`
}

type N01curd__create__createInternalInput struct {
	Email string `json:"email,omitempty"`
}

type N01curd__create__createResponseData struct {
	Blog_createOneUser N01curd__create__createResponseData_blog_createOneUser `json:"blog_createOneUser,omitempty"`
}

type N01curd__create__createResponseData_blog_createOneUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N01curd__create__createWithPostInput struct {
	Email string `json:"email,omitempty"`
}

type N01curd__create__createWithPostInternalInput struct {
	Email string `json:"email,omitempty"`
}

type N01curd__create__createWithPostResponseData struct {
	Blog_createOneUser N01curd__create__createWithPostResponseData_blog_createOneUser `json:"blog_createOneUser,omitempty"`
}

type N01curd__create__createWithPostResponseData_blog_createOneUser struct {
	Post []N01curd__create__createWithPostResponseData_blog_createOneUser_Post `json:"Post,omitempty"`
	Id   int64                                                                 `json:"id"`
}

type N01curd__create__createWithPostResponseData_blog_createOneUser_Post struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N01curd__delete__deleteAllInput struct {
}

type N01curd__delete__deleteAllInternalInput struct {
}

type N01curd__delete__deleteAllResponseData struct {
	Blog_deleteManyUser N01curd__delete__deleteAllResponseData_blog_deleteManyUser `json:"blog_deleteManyUser,omitempty"`
}

type N01curd__delete__deleteAllResponseData_blog_deleteManyUser struct {
	Count int64 `json:"count"`
}

type N01curd__delete__deleteInput struct {
	Email string `json:"email,omitempty"`
}

type N01curd__delete__deleteInternalInput struct {
	Email string `json:"email,omitempty"`
}

type N01curd__delete__deleteManyInput struct {
	Contains string `json:"contains,omitempty"`
}

type N01curd__delete__deleteManyInternalInput struct {
	Contains string `json:"contains,omitempty"`
}

type N01curd__delete__deleteManyResponseData struct {
	Blog_deleteManyUser N01curd__delete__deleteManyResponseData_blog_deleteManyUser `json:"blog_deleteManyUser,omitempty"`
}

type N01curd__delete__deleteManyResponseData_blog_deleteManyUser struct {
	Count int64 `json:"count"`
}

type N01curd__delete__deleteResponseData struct {
	Blog_deleteOneUser N01curd__delete__deleteResponseData_blog_deleteOneUser `json:"blog_deleteOneUser,omitempty"`
}

type N01curd__delete__deleteResponseData_blog_deleteOneUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N01curd__read__filterMultipleInput struct {
}

type N01curd__read__filterMultipleInternalInput struct {
}

type N01curd__read__filterMultipleResponseData struct {
	Blog_findManyUser []N01curd__read__filterMultipleResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N01curd__read__filterMultipleResponseData_blog_findManyUser struct {
	Post    []N01curd__read__filterMultipleResponseData_blog_findManyUser_Post  `json:"Post,omitempty"`
	Profile N01curd__read__filterMultipleResponseData_blog_findManyUser_Profile `json:"Profile,omitempty"`
	Email   string                                                              `json:"email"`
	Id      int64                                                               `json:"id"`
}

type N01curd__read__filterMultipleResponseData_blog_findManyUser_Post struct {
	Id        int64  `json:"id"`
	Like      int64  `json:"like,omitempty"`
	Published bool   `json:"published"`
	Title     string `json:"title"`
}

type N01curd__read__filterMultipleResponseData_blog_findManyUser_Profile struct {
	Address string `json:"address"`
}

type N01curd__read__filterSingInput struct {
	EndsWith string `json:"endsWith,omitempty"`
}

type N01curd__read__filterSingInternalInput struct {
	EndsWith string `json:"endsWith,omitempty"`
}

type N01curd__read__filterSingResponseData struct {
	Blog_findManyUser []N01curd__read__filterSingResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N01curd__read__filterSingResponseData_blog_findManyUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N01curd__read__findAllInput struct {
}

type N01curd__read__findAllInternalInput struct {
}

type N01curd__read__findAllResponseData struct {
	Blog_findManyUser []N01curd__read__findAllResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N01curd__read__findAllResponseData_blog_findManyUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N01curd__read__findFirstInput struct {
}

type N01curd__read__findFirstInternalInput struct {
}

type N01curd__read__findFirstResponseData struct {
	Blog_findFirstUser N01curd__read__findFirstResponseData_blog_findFirstUser `json:"blog_findFirstUser,omitempty"`
}

type N01curd__read__findFirstResponseData_blog_findFirstUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N01curd__read__findUniqueInput struct {
	Email string `json:"email,omitempty"`
}

type N01curd__read__findUniqueInternalInput struct {
	Email string `json:"email,omitempty"`
}

type N01curd__read__findUniqueResponseData struct {
	Blog_findUniqueUser N01curd__read__findUniqueResponseData_blog_findUniqueUser `json:"blog_findUniqueUser,omitempty"`
}

type N01curd__read__findUniqueResponseData_blog_findUniqueUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N01curd__update__incrementInput struct {
	Contains string `json:"contains,omitempty"`
}

type N01curd__update__incrementInternalInput struct {
	Contains string `json:"contains,omitempty"`
}

type N01curd__update__incrementResponseData struct {
	Blog_updateManyPost N01curd__update__incrementResponseData_blog_updateManyPost `json:"blog_updateManyPost,omitempty"`
}

type N01curd__update__incrementResponseData_blog_updateManyPost struct {
	Count int64 `json:"count"`
}

type N01curd__update__updateInput struct {
	Email string `json:"email,omitempty"`
	Id    int64  `json:"id,omitempty"`
}

type N01curd__update__updateInternalInput struct {
	Email string `json:"email,omitempty"`
	Id    int64  `json:"id,omitempty"`
}

type N01curd__update__updateManyInput struct {
	Email    string `json:"email,omitempty"`
	Newemail string `json:"newemail,omitempty"`
}

type N01curd__update__updateManyInternalInput struct {
	Email    string `json:"email,omitempty"`
	Newemail string `json:"newemail,omitempty"`
}

type N01curd__update__updateManyResponseData struct {
	Blog_updateManyUser N01curd__update__updateManyResponseData_blog_updateManyUser `json:"blog_updateManyUser,omitempty"`
}

type N01curd__update__updateManyResponseData_blog_updateManyUser struct {
	Count int64 `json:"count"`
}

type N01curd__update__updateResponseData struct {
	Blog_updateOneUser N01curd__update__updateResponseData_blog_updateOneUser `json:"blog_updateOneUser,omitempty"`
}

type N01curd__update__updateResponseData_blog_updateOneUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N01curd__update__upsertInput struct {
	Email string `json:"email,omitempty"`
}

type N01curd__update__upsertInternalInput struct {
	Email string `json:"email,omitempty"`
}

type N01curd__update__upsertResponseData struct {
	Blog_upsertOneUser N01curd__update__upsertResponseData_blog_upsertOneUser `json:"blog_upsertOneUser,omitempty"`
}

type N01curd__update__upsertResponseData_blog_upsertOneUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N02relation__filterListInput struct {
}

type N02relation__filterListInternalInput struct {
}

type N02relation__filterListResponseData struct {
	Blog_findFirstUser N02relation__filterListResponseData_blog_findFirstUser `json:"blog_findFirstUser,omitempty"`
}

type N02relation__filterListResponseData_blog_findFirstUser struct {
	Post  []N02relation__filterListResponseData_blog_findFirstUser_Post `json:"Post,omitempty"`
	Email string                                                        `json:"email"`
	Id    int64                                                         `json:"id"`
}

type N02relation__filterListResponseData_blog_findFirstUser_Post struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N02relation__filter__filter2ManyInput struct {
}

type N02relation__filter__filter2ManyInternalInput struct {
}

type N02relation__filter__filter2ManyNoneInput struct {
}

type N02relation__filter__filter2ManyNoneInternalInput struct {
}

type N02relation__filter__filter2ManyNoneResponseData struct {
	Blog_findManyUser []N02relation__filter__filter2ManyNoneResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N02relation__filter__filter2ManyNoneResponseData_blog_findManyUser struct {
	Post  []N02relation__filter__filter2ManyNoneResponseData_blog_findManyUser_Post `json:"Post,omitempty"`
	Email string                                                                    `json:"email"`
	Id    int64                                                                     `json:"id"`
}

type N02relation__filter__filter2ManyNoneResponseData_blog_findManyUser_Post struct {
	Id   int64 `json:"id"`
	Like int64 `json:"like,omitempty"`
}

type N02relation__filter__filter2ManyResponseData struct {
	Blog_findManyUser []N02relation__filter__filter2ManyResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N02relation__filter__filter2ManyResponseData_blog_findManyUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N02relation__filter__filter2ManySomeInput struct {
	EndsWith string `json:"endsWith,omitempty"`
	Equals   bool   `json:"equals,omitempty"`
}

type N02relation__filter__filter2ManySomeInternalInput struct {
	EndsWith string `json:"endsWith,omitempty"`
	Equals   bool   `json:"equals,omitempty"`
}

type N02relation__filter__filter2ManySomeResponseData struct {
	Blog_findManyUser []N02relation__filter__filter2ManySomeResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N02relation__filter__filter2ManySomeResponseData_blog_findManyUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N02relation__filter__filter2One2Input struct {
}

type N02relation__filter__filter2One2InternalInput struct {
}

type N02relation__filter__filter2One2ResponseData struct {
	Blog_findManyPost []N02relation__filter__filter2One2ResponseData_blog_findManyPost `json:"blog_findManyPost"`
}

type N02relation__filter__filter2One2ResponseData_blog_findManyPost struct {
	User  N02relation__filter__filter2One2ResponseData_blog_findManyPost_User `json:"User,omitempty"`
	Id    int64                                                               `json:"id"`
	Title string                                                              `json:"title"`
}

type N02relation__filter__filter2One2ResponseData_blog_findManyPost_User struct {
	Id int64 `json:"id"`
}

type N02relation__filter__filter2OneInput struct {
}

type N02relation__filter__filter2OneInternalInput struct {
}

type N02relation__filter__filter2OneNoneInput struct {
}

type N02relation__filter__filter2OneNoneInternalInput struct {
}

type N02relation__filter__filter2OneNoneResponseData struct {
	Blog_findManyPost []N02relation__filter__filter2OneNoneResponseData_blog_findManyPost `json:"blog_findManyPost"`
}

type N02relation__filter__filter2OneNoneResponseData_blog_findManyPost struct {
	User      N02relation__filter__filter2OneNoneResponseData_blog_findManyPost_User `json:"User,omitempty"`
	Id        int64                                                                  `json:"id"`
	Like      int64                                                                  `json:"like,omitempty"`
	Published bool                                                                   `json:"published"`
	Title     string                                                                 `json:"title"`
}

type N02relation__filter__filter2OneNoneResponseData_blog_findManyPost_User struct {
	Id int64 `json:"id"`
}

type N02relation__filter__filter2OneResponseData struct {
	Blog_findManyPost []N02relation__filter__filter2OneResponseData_blog_findManyPost `json:"blog_findManyPost"`
}

type N02relation__filter__filter2OneResponseData_blog_findManyPost struct {
	Id int64 `json:"id"`
}

type N02relation__filter__filterRelateSomeInput struct {
}

type N02relation__filter__filterRelateSomeInternalInput struct {
}

type N02relation__filter__filterRelateSomeResponseData struct {
	Blog_findManyUser []N02relation__filter__filterRelateSomeResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N02relation__filter__filterRelateSomeResponseData_blog_findManyUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N02relation__nestedReadInput struct {
	Email string `json:"email,omitempty"`
}

type N02relation__nestedReadInternalInput struct {
	Email string `json:"email,omitempty"`
}

type N02relation__nestedReadResponseData struct {
	Blog_findUniqueUser N02relation__nestedReadResponseData_blog_findUniqueUser `json:"blog_findUniqueUser,omitempty"`
}

type N02relation__nestedReadResponseData_blog_findUniqueUser struct {
	Post    []N02relation__nestedReadResponseData_blog_findUniqueUser_Post  `json:"Post,omitempty"`
	Profile N02relation__nestedReadResponseData_blog_findUniqueUser_Profile `json:"Profile,omitempty"`
	Email   string                                                          `json:"email"`
	Id      int64                                                           `json:"id"`
}

type N02relation__nestedReadResponseData_blog_findUniqueUser_Post struct {
	Category  []N02relation__nestedReadResponseData_blog_findUniqueUser_Post_Category `json:"Category,omitempty"`
	AuthorId  int64                                                                   `json:"authorId,omitempty"`
	Id        int64                                                                   `json:"id"`
	Published bool                                                                    `json:"published"`
	Title     string                                                                  `json:"title"`
}

type N02relation__nestedReadResponseData_blog_findUniqueUser_Post_Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type N02relation__nestedReadResponseData_blog_findUniqueUser_Profile struct {
	Address string `json:"address"`
	Id      int64  `json:"id"`
	UserId  int64  `json:"userId"`
}

type N02relation__realationCountInput struct {
	Email string `json:"email,omitempty"`
}

type N02relation__realationCountInternalInput struct {
	Email string `json:"email,omitempty"`
}

type N02relation__realationCountResponseData struct {
	Blog_findUniqueUser N02relation__realationCountResponseData_blog_findUniqueUser `json:"blog_findUniqueUser,omitempty"`
}

type N02relation__realationCountResponseData_blog_findUniqueUser struct {
	Email     string                                                                `json:"email"`
	Id        int64                                                                 `json:"id"`
	TotalPost N02relation__realationCountResponseData_blog_findUniqueUser_totalPost `json:"totalPost"`
}

type N02relation__realationCountResponseData_blog_findUniqueUser_totalPost struct {
	Post int64 `json:"Post"`
}

type N02relation__write__conectSingRecordsInput struct {
}

type N02relation__write__conectSingRecordsInternalInput struct {
}

type N02relation__write__conectSingRecordsResponseData struct {
	Blog_createOneUser N02relation__write__conectSingRecordsResponseData_blog_createOneUser `json:"blog_createOneUser,omitempty"`
}

type N02relation__write__conectSingRecordsResponseData_blog_createOneUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N02relation__write__connectInput struct {
	Id int64 `json:"id,omitempty"`
}

type N02relation__write__connectInternalInput struct {
	Id int64 `json:"id,omitempty"`
}

type N02relation__write__connectOrCreateRecordInput struct {
	Email string `json:"email,omitempty"`
}

type N02relation__write__connectOrCreateRecordInternalInput struct {
	Email string `json:"email,omitempty"`
}

type N02relation__write__connectOrCreateRecordResponseData struct {
	Blog_createOnePost N02relation__write__connectOrCreateRecordResponseData_blog_createOnePost `json:"blog_createOnePost,omitempty"`
}

type N02relation__write__connectOrCreateRecordResponseData_blog_createOnePost struct {
	Id int64 `json:"id"`
}

type N02relation__write__connectResponseData struct {
	Blog_updateOnePost N02relation__write__connectResponseData_blog_updateOnePost `json:"blog_updateOnePost,omitempty"`
}

type N02relation__write__connectResponseData_blog_updateOnePost struct {
	Id int64 `json:"id"`
}

type N02relation__write__deleteAllRelatedRecordsInput struct {
	Id int64 `json:"id,omitempty"`
}

type N02relation__write__deleteAllRelatedRecordsInternalInput struct {
	Id int64 `json:"id,omitempty"`
}

type N02relation__write__deleteAllRelatedRecordsResponseData struct {
	Blog_updateOneUser N02relation__write__deleteAllRelatedRecordsResponseData_blog_updateOneUser `json:"blog_updateOneUser,omitempty"`
}

type N02relation__write__deleteAllRelatedRecordsResponseData_blog_updateOneUser struct {
	Id int64 `json:"id"`
}

type N02relation__write__deleteSpecificRelatedRecords2Input struct {
	Id int64 `json:"id,omitempty"`
}

type N02relation__write__deleteSpecificRelatedRecords2InternalInput struct {
	Id int64 `json:"id,omitempty"`
}

type N02relation__write__deleteSpecificRelatedRecords2ResponseData struct {
	Blog_updateOneUser N02relation__write__deleteSpecificRelatedRecords2ResponseData_blog_updateOneUser `json:"blog_updateOneUser,omitempty"`
}

type N02relation__write__deleteSpecificRelatedRecords2ResponseData_blog_updateOneUser struct {
	Count N02relation__write__deleteSpecificRelatedRecords2ResponseData_blog_updateOneUser__count `json:"_count"`
	Id    int64                                                                                   `json:"id"`
}

type N02relation__write__deleteSpecificRelatedRecords2ResponseData_blog_updateOneUser__count struct {
	Post int64 `json:"Post"`
}

type N02relation__write__deleteSpecificRelatedRecordsInput struct {
	Id int64 `json:"id,omitempty"`
}

type N02relation__write__deleteSpecificRelatedRecordsInternalInput struct {
	Id int64 `json:"id,omitempty"`
}

type N02relation__write__deleteSpecificRelatedRecordsResponseData struct {
	Blog_updateOneUser N02relation__write__deleteSpecificRelatedRecordsResponseData_blog_updateOneUser `json:"blog_updateOneUser,omitempty"`
}

type N02relation__write__deleteSpecificRelatedRecordsResponseData_blog_updateOneUser struct {
	Id int64 `json:"id"`
}

type N02relation__write__disconnectInput struct {
	Id int64 `json:"id,omitempty"`
}

type N02relation__write__disconnectInternalInput struct {
	Id int64 `json:"id,omitempty"`
}

type N02relation__write__disconnectResponseData struct {
	Blog_updateOnePost N02relation__write__disconnectResponseData_blog_updateOnePost `json:"blog_updateOnePost,omitempty"`
}

type N02relation__write__disconnectResponseData_blog_updateOnePost struct {
	Id int64 `json:"id"`
}

type N02relation__write__updateAllRelatedRecordsInput struct {
}

type N02relation__write__updateAllRelatedRecordsInternalInput struct {
}

type N02relation__write__updateAllRelatedRecordsResponseData struct {
	Blog_updateOneUser N02relation__write__updateAllRelatedRecordsResponseData_blog_updateOneUser `json:"blog_updateOneUser,omitempty"`
}

type N02relation__write__updateAllRelatedRecordsResponseData_blog_updateOneUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N02relation__write__updateOrCreateRelatedRecordInput struct {
	Email string `json:"email,omitempty"`
}

type N02relation__write__updateOrCreateRelatedRecordInternalInput struct {
	Email string `json:"email,omitempty"`
}

type N02relation__write__updateOrCreateRelatedRecordResponseData struct {
	Blog_updateOnePost N02relation__write__updateOrCreateRelatedRecordResponseData_blog_updateOnePost `json:"blog_updateOnePost,omitempty"`
}

type N02relation__write__updateOrCreateRelatedRecordResponseData_blog_updateOnePost struct {
	Id int64 `json:"id"`
}

type N02relation__write__updateSpecificRelatedRecordInput struct {
}

type N02relation__write__updateSpecificRelatedRecordInternalInput struct {
}

type N02relation__write__updateSpecificRelatedRecordResponseData struct {
	Blog_updateOneUser N02relation__write__updateSpecificRelatedRecordResponseData_blog_updateOneUser `json:"blog_updateOneUser,omitempty"`
}

type N02relation__write__updateSpecificRelatedRecordResponseData_blog_updateOneUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N03filterAndSorting__filter__combiningInput struct {
}

type N03filterAndSorting__filter__combiningInternalInput struct {
}

type N03filterAndSorting__filter__combiningResponseData struct {
	Blog_findManyUser []N03filterAndSorting__filter__combiningResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N03filterAndSorting__filter__combiningResponseData_blog_findManyUser struct {
	Email string `json:"email"`
}

type N03filterAndSorting__filter__filterInput struct {
}

type N03filterAndSorting__filter__filterInternalInput struct {
}

type N03filterAndSorting__filter__filterResponseData struct {
	Blog_findManyUser []N03filterAndSorting__filter__filterResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N03filterAndSorting__filter__filterResponseData_blog_findManyUser struct {
	Post  []N03filterAndSorting__filter__filterResponseData_blog_findManyUser_Post `json:"Post,omitempty"`
	Email string                                                                   `json:"email"`
	Id    int64                                                                    `json:"id"`
}

type N03filterAndSorting__filter__filterResponseData_blog_findManyUser_Post struct {
	Id        int64  `json:"id"`
	Published bool   `json:"published"`
	Title     string `json:"title"`
}

type N03filterAndSorting__filter__nonnullInput struct {
}

type N03filterAndSorting__filter__nonnullInternalInput struct {
}

type N03filterAndSorting__filter__nonnullResponseData struct {
	Blog_findManyPost []N03filterAndSorting__filter__nonnullResponseData_blog_findManyPost `json:"blog_findManyPost"`
}

type N03filterAndSorting__filter__nonnullResponseData_blog_findManyPost struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N03filterAndSorting__filter__nullInput struct {
}

type N03filterAndSorting__filter__nullInternalInput struct {
}

type N03filterAndSorting__filter__nullResponseData struct {
	Blog_findManyPost []N03filterAndSorting__filter__nullResponseData_blog_findManyPost `json:"blog_findManyPost"`
}

type N03filterAndSorting__filter__nullResponseData_blog_findManyPost struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N03filterAndSorting__filter__pmsymbolsInput struct {
}

type N03filterAndSorting__filter__pmsymbolsInternalInput struct {
}

type N03filterAndSorting__filter__pmsymbolsResponseData struct {
	Blog_findManyUser []N03filterAndSorting__filter__pmsymbolsResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N03filterAndSorting__filter__pmsymbolsResponseData_blog_findManyUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N03filterAndSorting__sort__relationAggInput struct {
}

type N03filterAndSorting__sort__relationAggInternalInput struct {
}

type N03filterAndSorting__sort__relationAggResponseData struct {
	Blog_findManyUser []N03filterAndSorting__sort__relationAggResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N03filterAndSorting__sort__relationAggResponseData_blog_findManyUser struct {
	Count N03filterAndSorting__sort__relationAggResponseData_blog_findManyUser__count `json:"_count"`
	Email string                                                                      `json:"email"`
	Id    int64                                                                       `json:"id"`
}

type N03filterAndSorting__sort__relationAggResponseData_blog_findManyUser__count struct {
	Post int64 `json:"Post"`
}

type N03filterAndSorting__sort__relationInput struct {
}

type N03filterAndSorting__sort__relationInternalInput struct {
}

type N03filterAndSorting__sort__relationResponseData struct {
	Blog_findManyPost []N03filterAndSorting__sort__relationResponseData_blog_findManyPost `json:"blog_findManyPost"`
}

type N03filterAndSorting__sort__relationResponseData_blog_findManyPost struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N03filterAndSorting__sort__sortInput struct {
}

type N03filterAndSorting__sort__sortInternalInput struct {
}

type N03filterAndSorting__sort__sortResponseData struct {
	Blog_findManyUser []N03filterAndSorting__sort__sortResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N03filterAndSorting__sort__sortResponseData_blog_findManyUser struct {
	Post  []N03filterAndSorting__sort__sortResponseData_blog_findManyUser_Post `json:"Post,omitempty"`
	Email string                                                               `json:"email"`
	Id    int64                                                                `json:"id"`
}

type N03filterAndSorting__sort__sortResponseData_blog_findManyUser_Post struct {
	Content string `json:"content,omitempty"`
	Id      int64  `json:"id"`
	Title   string `json:"title"`
}

type N04pagination__cursorBackwardsInput struct {
}

type N04pagination__cursorBackwardsInternalInput struct {
}

type N04pagination__cursorBackwardsResponseData struct {
	Blog_findManyPost []N04pagination__cursorBackwardsResponseData_blog_findManyPost `json:"blog_findManyPost"`
}

type N04pagination__cursorBackwardsResponseData_blog_findManyPost struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N04pagination__cursorInput struct {
}

type N04pagination__cursorInternalInput struct {
}

type N04pagination__cursorResponseData struct {
	Blog_findManyPost []N04pagination__cursorResponseData_blog_findManyPost `json:"blog_findManyPost"`
}

type N04pagination__cursorResponseData_blog_findManyPost struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N04pagination__filteringAndOffsetInput struct {
}

type N04pagination__filteringAndOffsetInternalInput struct {
}

type N04pagination__filteringAndOffsetResponseData struct {
	Blog_findManyPost []N04pagination__filteringAndOffsetResponseData_blog_findManyPost `json:"blog_findManyPost"`
}

type N04pagination__filteringAndOffsetResponseData_blog_findManyPost struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N04pagination__offsetInput struct {
}

type N04pagination__offsetInternalInput struct {
}

type N04pagination__offsetResponseData struct {
	Blog_findManyPost []N04pagination__offsetResponseData_blog_findManyPost `json:"blog_findManyPost"`
}

type N04pagination__offsetResponseData_blog_findManyPost struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N04pagination__sortingAndOffsetInput struct {
}

type N04pagination__sortingAndOffsetInternalInput struct {
}

type N04pagination__sortingAndOffsetResponseData struct {
	Blog_findManyPost []N04pagination__sortingAndOffsetResponseData_blog_findManyPost `json:"blog_findManyPost"`
}

type N04pagination__sortingAndOffsetResponseData_blog_findManyPost struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N05aggregation__aggregate__avgInput struct {
}

type N05aggregation__aggregate__avgInternalInput struct {
}

type N05aggregation__aggregate__avgResponseData struct {
	Blog_aggregateUser N05aggregation__aggregate__avgResponseData_blog_aggregateUser `json:"blog_aggregateUser"`
}

type N05aggregation__aggregate__avgResponseData_blog_aggregateUser struct {
	Avg N05aggregation__aggregate__avgResponseData_blog_aggregateUser__avg `json:"_avg,omitempty"`
}

type N05aggregation__aggregate__avgResponseData_blog_aggregateUser__avg struct {
	Age float64 `json:"age,omitempty"`
}

type N05aggregation__aggregate__nullableInput struct {
}

type N05aggregation__aggregate__nullableInternalInput struct {
}

type N05aggregation__aggregate__nullableResponseData struct {
	Blog_aggregateUser N05aggregation__aggregate__nullableResponseData_blog_aggregateUser `json:"blog_aggregateUser"`
}

type N05aggregation__aggregate__nullableResponseData_blog_aggregateUser struct {
	Avg   N05aggregation__aggregate__nullableResponseData_blog_aggregateUser__avg   `json:"_avg,omitempty"`
	Count N05aggregation__aggregate__nullableResponseData_blog_aggregateUser__count `json:"_count,omitempty"`
}

type N05aggregation__aggregate__nullableResponseData_blog_aggregateUser__avg struct {
	Age float64 `json:"age,omitempty"`
}

type N05aggregation__aggregate__nullableResponseData_blog_aggregateUser__count struct {
	Age int64 `json:"age"`
	Id  int64 `json:"id"`
}

type N05aggregation__count__countInput struct {
}

type N05aggregation__count__countInternalInput struct {
}

type N05aggregation__count__countResponseData struct {
	Blog_aggregateUser N05aggregation__count__countResponseData_blog_aggregateUser `json:"blog_aggregateUser"`
}

type N05aggregation__count__countResponseData_blog_aggregateUser struct {
	Count N05aggregation__count__countResponseData_blog_aggregateUser__count `json:"_count,omitempty"`
}

type N05aggregation__count__countResponseData_blog_aggregateUser__count struct {
	Id int64 `json:"id"`
}

type N05aggregation__count__filterInput struct {
}

type N05aggregation__count__filterInternalInput struct {
}

type N05aggregation__count__filterResponseData struct {
	Blog_aggregateUser N05aggregation__count__filterResponseData_blog_aggregateUser `json:"blog_aggregateUser"`
}

type N05aggregation__count__filterResponseData_blog_aggregateUser struct {
	Count N05aggregation__count__filterResponseData_blog_aggregateUser__count `json:"_count,omitempty"`
}

type N05aggregation__count__filterResponseData_blog_aggregateUser__count struct {
	All int64 `json:"_all"`
}

type N05aggregation__count__nonnullFieldInput struct {
}

type N05aggregation__count__nonnullFieldInternalInput struct {
}

type N05aggregation__count__nonnullFieldResponseData struct {
	Blog_aggregateUser N05aggregation__count__nonnullFieldResponseData_blog_aggregateUser `json:"blog_aggregateUser"`
}

type N05aggregation__count__nonnullFieldResponseData_blog_aggregateUser struct {
	Count N05aggregation__count__nonnullFieldResponseData_blog_aggregateUser__count `json:"_count,omitempty"`
}

type N05aggregation__count__nonnullFieldResponseData_blog_aggregateUser__count struct {
	All     int64 `json:"_all"`
	Country int64 `json:"country"`
	Email   int64 `json:"email"`
}

type N05aggregation__count__relationFilterInput struct {
}

type N05aggregation__count__relationFilterInternalInput struct {
}

type N05aggregation__count__relationFilterResponseData struct {
	Blog_findManyUser []N05aggregation__count__relationFilterResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N05aggregation__count__relationFilterResponseData_blog_findManyUser struct {
	Count N05aggregation__count__relationFilterResponseData_blog_findManyUser__count `json:"_count"`
	Email string                                                                     `json:"email"`
	Id    int64                                                                      `json:"id"`
}

type N05aggregation__count__relationFilterResponseData_blog_findManyUser__count struct {
	Post int64 `json:"Post"`
}

type N05aggregation__count__relationInput struct {
}

type N05aggregation__count__relationInternalInput struct {
}

type N05aggregation__count__relationResponseData struct {
	Blog_findManyUser []N05aggregation__count__relationResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N05aggregation__count__relationResponseData_blog_findManyUser struct {
	Count N05aggregation__count__relationResponseData_blog_findManyUser__count `json:"_count"`
	Email string                                                               `json:"email"`
	Id    int64                                                                `json:"id"`
}

type N05aggregation__count__relationResponseData_blog_findManyUser__count struct {
	Post int64 `json:"Post"`
}

type N05aggregation__distinct__distinctInput struct {
}

type N05aggregation__distinct__distinctInternalInput struct {
}

type N05aggregation__distinct__distinctResponseData struct {
	Blog_findManyUser []N05aggregation__distinct__distinctResponseData_blog_findManyUser `json:"blog_findManyUser"`
}

type N05aggregation__distinct__distinctResponseData_blog_findManyUser struct {
	Country string `json:"country,omitempty"`
}

type N05aggregation__group__groupFilterInput struct {
}

type N05aggregation__group__groupFilterInternalInput struct {
}

type N05aggregation__group__groupFilterResponseData struct {
	Blog_groupByUser []N05aggregation__group__groupFilterResponseData_blog_groupByUser `json:"blog_groupByUser"`
}

type N05aggregation__group__groupFilterResponseData_blog_groupByUser struct {
	Sum     N05aggregation__group__groupFilterResponseData_blog_groupByUser__sum `json:"_sum,omitempty"`
	Country string                                                               `json:"country,omitempty"`
}

type N05aggregation__group__groupFilterResponseData_blog_groupByUser__sum struct {
	Age int64 `json:"age,omitempty"`
	Id  int64 `json:"id,omitempty"`
}

type N05aggregation__group__groupHavingInput struct {
}

type N05aggregation__group__groupHavingInternalInput struct {
}

type N05aggregation__group__groupHavingResponseData struct {
	Blog_groupByUser []N05aggregation__group__groupHavingResponseData_blog_groupByUser `json:"blog_groupByUser"`
}

type N05aggregation__group__groupHavingResponseData_blog_groupByUser struct {
	Sum     N05aggregation__group__groupHavingResponseData_blog_groupByUser__sum `json:"_sum,omitempty"`
	Country string                                                               `json:"country,omitempty"`
}

type N05aggregation__group__groupHavingResponseData_blog_groupByUser__sum struct {
	Age int64 `json:"age,omitempty"`
	Id  int64 `json:"id,omitempty"`
}

type N05aggregation__group__groupInput struct {
}

type N05aggregation__group__groupInternalInput struct {
}

type N05aggregation__group__groupOrderbyFieldInput struct {
}

type N05aggregation__group__groupOrderbyFieldInternalInput struct {
}

type N05aggregation__group__groupOrderbyFieldResponseData struct {
	Blog_groupByUser []N05aggregation__group__groupOrderbyFieldResponseData_blog_groupByUser `json:"blog_groupByUser"`
}

type N05aggregation__group__groupOrderbyFieldResponseData_blog_groupByUser struct {
	Count   N05aggregation__group__groupOrderbyFieldResponseData_blog_groupByUser__count `json:"_count,omitempty"`
	Country string                                                                       `json:"country,omitempty"`
}

type N05aggregation__group__groupOrderbyFieldResponseData_blog_groupByUser__count struct {
	Country int64 `json:"country"`
}

type N05aggregation__group__groupOrderbyInput struct {
}

type N05aggregation__group__groupOrderbyInternalInput struct {
}

type N05aggregation__group__groupOrderbyResponseData struct {
	Blog_groupByUser []N05aggregation__group__groupOrderbyResponseData_blog_groupByUser `json:"blog_groupByUser"`
}

type N05aggregation__group__groupOrderbyResponseData_blog_groupByUser struct {
	Count   N05aggregation__group__groupOrderbyResponseData_blog_groupByUser__count `json:"_count,omitempty"`
	Country string                                                                  `json:"country,omitempty"`
}

type N05aggregation__group__groupOrderbyResponseData_blog_groupByUser__count struct {
	Country int64 `json:"country"`
}

type N05aggregation__group__groupResponseData struct {
	Blog_groupByUser []N05aggregation__group__groupResponseData_blog_groupByUser `json:"blog_groupByUser"`
}

type N05aggregation__group__groupResponseData_blog_groupByUser struct {
	Sum     N05aggregation__group__groupResponseData_blog_groupByUser__sum `json:"_sum,omitempty"`
	Country string                                                         `json:"country,omitempty"`
}

type N05aggregation__group__groupResponseData_blog_groupByUser__sum struct {
	Age int64 `json:"age,omitempty"`
	Id  int64 `json:"id,omitempty"`
}

type N06advanced__join__mutationInput struct {
}

type N06advanced__join__mutationInternalInput struct {
}

type N06advanced__join__mutationResponseData struct {
	Blog_findUniqueUser N06advanced__join__mutationResponseData_blog_findUniqueUser `json:"blog_findUniqueUser,omitempty"`
}

type N06advanced__join__mutationResponseData_blog_findUniqueUser struct {
	Join_mutation N06advanced__join__mutationResponseData_blog_findUniqueUser__join_mutation `json:"_join_mutation"`
	Email         string                                                                     `json:"email"`
	Id            int64                                                                      `json:"id"`
}

type N06advanced__join__mutationResponseData_blog_findUniqueUser__join_mutation struct {
	Todo_updateManyTodo N06advanced__join__mutationResponseData_blog_findUniqueUser__join_mutation_todo_updateManyTodo `json:"todo_updateManyTodo,omitempty"`
}

type N06advanced__join__mutationResponseData_blog_findUniqueUser__join_mutation_todo_updateManyTodo struct {
	Count int64 `json:"count"`
}

type N06advanced__join__queryInput struct {
}

type N06advanced__join__queryInternalInput struct {
}

type N06advanced__join__queryResponseData struct {
	Blog_findUniqueUser N06advanced__join__queryResponseData_blog_findUniqueUser `json:"blog_findUniqueUser,omitempty"`
}

type N06advanced__join__queryResponseData_blog_findUniqueUser struct {
	Join  N06advanced__join__queryResponseData_blog_findUniqueUser__join `json:"_join"`
	Email string                                                         `json:"email"`
	Id    int64                                                          `json:"id"`
}

type N06advanced__join__queryResponseData_blog_findUniqueUser__join struct {
	Todo_findManyTodo []N06advanced__join__queryResponseData_blog_findUniqueUser__join_todo_findManyTodo `json:"todo_findManyTodo"`
}

type N06advanced__join__queryResponseData_blog_findUniqueUser__join_todo_findManyTodo struct {
	Completed bool   `json:"completed"`
	Id        int64  `json:"id"`
	Title     string `json:"title"`
}

type N06advanced__raw__executeInput struct {
	Parameters any `json:"parameters,omitempty"`
}

type N06advanced__raw__executeInternalInput struct {
	Parameters any `json:"parameters,omitempty"`
}

type N06advanced__raw__executeResponseData struct {
	Blog_executeRaw any `json:"blog_executeRaw,omitempty"`
}

type N06advanced__raw__queryInput struct {
	Parameters any `json:"parameters,omitempty"`
}

type N06advanced__raw__queryInternalInput struct {
	Parameters any `json:"parameters,omitempty"`
}

type N06advanced__raw__queryResponseData struct {
	Blog_queryRaw any `json:"blog_queryRaw,omitempty"`
}

type N06advanced__selection_pro__transformInput struct {
	Email string `json:"email,omitempty"`
}

type N06advanced__selection_pro__transformInternalInput struct {
	Email string `json:"email,omitempty"`
}

type N06advanced__selection_pro__transformResponseData struct {
	Blog_findUniqueUser N06advanced__selection_pro__transformResponseData_blog_findUniqueUser `json:"blog_findUniqueUser,omitempty"`
}

type N06advanced__selection_pro__transformResponseData_blog_findUniqueUser struct {
	Country   string   `json:"country,omitempty"`
	Email     string   `json:"email"`
	Id        int64    `json:"id"`
	PostCount int64    `json:"postCount"`
	Posts     []string `json:"posts,omitempty"`
}

type N06advanced__setting__livequeryInput struct {
	Uid int64 `json:"uid,omitempty"`
}

type N06advanced__setting__livequeryInternalInput struct {
	Uid int64 `json:"uid,omitempty"`
}

type N06advanced__setting__livequeryResponseData struct {
	Blog_aggregatePost N06advanced__setting__livequeryResponseData_blog_aggregatePost `json:"blog_aggregatePost"`
}

type N06advanced__setting__livequeryResponseData_blog_aggregatePost struct {
	Count N06advanced__setting__livequeryResponseData_blog_aggregatePost__count `json:"_count,omitempty"`
}

type N06advanced__setting__livequeryResponseData_blog_aggregatePost__count struct {
	All int64 `json:"_all"`
}

type N06advanced__setting__rateLimitInput struct {
	Title string `json:"title,omitempty"`
}

type N06advanced__setting__rateLimitInternalInput struct {
	Title string `json:"title,omitempty"`
}

type N06advanced__setting__rateLimitResponseData struct {
	Blog_createOnePost N06advanced__setting__rateLimitResponseData_blog_createOnePost `json:"blog_createOnePost,omitempty"`
}

type N06advanced__setting__rateLimitResponseData_blog_createOnePost struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N06advanced__transaction__transactionDelInput struct {
	UserId int64 `json:"userId,omitempty"`
}

type N06advanced__transaction__transactionDelInternalInput struct {
	UserId int64 `json:"userId,omitempty"`
}

type N06advanced__transaction__transactionDelResponseData struct {
	Blog_deleteOneProfile N06advanced__transaction__transactionDelResponseData_blog_deleteOneProfile `json:"blog_deleteOneProfile,omitempty"`
	Blog_deleteOneUser    N06advanced__transaction__transactionDelResponseData_blog_deleteOneUser    `json:"blog_deleteOneUser,omitempty"`
}

type N06advanced__transaction__transactionDelResponseData_blog_deleteOneProfile struct {
	Id int64 `json:"id"`
}

type N06advanced__transaction__transactionDelResponseData_blog_deleteOneUser struct {
	Id int64 `json:"id"`
}

type N06advanced__variable_pro__fromheaderInput struct {
}

type N06advanced__variable_pro__fromheaderInternalInput struct {
	Title string `json:"title"`
}

type N06advanced__variable_pro__fromheaderResponseData struct {
	Todo_createOneTodo N06advanced__variable_pro__fromheaderResponseData_todo_createOneTodo `json:"todo_createOneTodo,omitempty"`
}

type N06advanced__variable_pro__fromheaderResponseData_todo_createOneTodo struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N06advanced__variable_pro__injectcurrentdatetimeInput struct {
}

type N06advanced__variable_pro__injectcurrentdatetimeInternalInput struct {
	CreatedAt time.Time `json:"createdAt"`
}

type N06advanced__variable_pro__injectcurrentdatetimeResponseData struct {
	Todo_createOneTodo N06advanced__variable_pro__injectcurrentdatetimeResponseData_todo_createOneTodo `json:"todo_createOneTodo,omitempty"`
}

type N06advanced__variable_pro__injectcurrentdatetimeResponseData_todo_createOneTodo struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
}

type N06advanced__variable_pro__injectcurrentdatetime_offsetInput struct {
}

type N06advanced__variable_pro__injectcurrentdatetime_offsetInternalInput struct {
	CreatedAt time.Time `json:"createdAt"`
}

type N06advanced__variable_pro__injectcurrentdatetime_offsetResponseData struct {
	Todo_createOneTodo N06advanced__variable_pro__injectcurrentdatetime_offsetResponseData_todo_createOneTodo `json:"todo_createOneTodo,omitempty"`
}

type N06advanced__variable_pro__injectcurrentdatetime_offsetResponseData_todo_createOneTodo struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
}

type N06advanced__variable_pro__injectenvironmentvariableInput struct {
}

type N06advanced__variable_pro__injectenvironmentvariableInternalInput struct {
	Title string `json:"title"`
}

type N06advanced__variable_pro__injectenvironmentvariableResponseData struct {
	Todo_createOneTodo N06advanced__variable_pro__injectenvironmentvariableResponseData_todo_createOneTodo `json:"todo_createOneTodo,omitempty"`
}

type N06advanced__variable_pro__injectenvironmentvariableResponseData_todo_createOneTodo struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N06advanced__variable_pro__injectgenerateduuidInput struct {
}

type N06advanced__variable_pro__injectgenerateduuidInternalInput struct {
	Title string `json:"title"`
}

type N06advanced__variable_pro__injectgenerateduuidResponseData struct {
	Todo_createOneTodo N06advanced__variable_pro__injectgenerateduuidResponseData_todo_createOneTodo `json:"todo_createOneTodo,omitempty"`
}

type N06advanced__variable_pro__injectgenerateduuidResponseData_todo_createOneTodo struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N06advanced__variable_pro__jsonschema__emailInput struct {
	Email string `json:"email"`
}

type N06advanced__variable_pro__jsonschema__emailInternalInput struct {
	Email string `json:"email"`
}

type N06advanced__variable_pro__jsonschema__emailResponseData struct {
	Blog_updateOneUser N06advanced__variable_pro__jsonschema__emailResponseData_blog_updateOneUser `json:"blog_updateOneUser,omitempty"`
}

type N06advanced__variable_pro__jsonschema__emailResponseData_blog_updateOneUser struct {
	Age int64 `json:"age,omitempty"`
	Id  int64 `json:"id"`
}

type N06advanced__variable_pro__jsonschema__numberInput struct {
	Age int64 `json:"age"`
}

type N06advanced__variable_pro__jsonschema__numberInternalInput struct {
	Age int64 `json:"age"`
}

type N06advanced__variable_pro__jsonschema__numberResponseData struct {
	Blog_updateOneUser N06advanced__variable_pro__jsonschema__numberResponseData_blog_updateOneUser `json:"blog_updateOneUser,omitempty"`
}

type N06advanced__variable_pro__jsonschema__numberResponseData_blog_updateOneUser struct {
	Age int64 `json:"age,omitempty"`
	Id  int64 `json:"id"`
}

type N06advanced__variable_pro__jsonschema__patternInput struct {
	Email string `json:"email"`
}

type N06advanced__variable_pro__jsonschema__patternInternalInput struct {
	Email string `json:"email"`
}

type N06advanced__variable_pro__jsonschema__patternResponseData struct {
	Blog_updateOneUser N06advanced__variable_pro__jsonschema__patternResponseData_blog_updateOneUser `json:"blog_updateOneUser,omitempty"`
}

type N06advanced__variable_pro__jsonschema__patternResponseData_blog_updateOneUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N06advanced__variable_pro__jsonschema__stringInput struct {
	Title string `json:"title"`
}

type N06advanced__variable_pro__jsonschema__stringInternalInput struct {
	Title string `json:"title"`
}

type N06advanced__variable_pro__jsonschema__stringResponseData struct {
	Todo_createOneTodo N06advanced__variable_pro__jsonschema__stringResponseData_todo_createOneTodo `json:"todo_createOneTodo,omitempty"`
}

type N06advanced__variable_pro__jsonschema__stringResponseData_todo_createOneTodo struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N07oidc__authInput struct {
}

type N07oidc__authInternalInput struct {
}

type N07oidc__authResponseData struct {
	Blog_deleteOneUser N07oidc__authResponseData_blog_deleteOneUser `json:"blog_deleteOneUser,omitempty"`
}

type N07oidc__authResponseData_blog_deleteOneUser struct {
	Id int64 `json:"id"`
}

type N07oidc__fromclaim__fromclaimInput struct {
}

type N07oidc__fromclaim__fromclaimInternalInput struct {
	Email string `json:"email"`
}

type N07oidc__fromclaim__fromclaimResponseData struct {
	Blog_findUniqueUser N07oidc__fromclaim__fromclaimResponseData_blog_findUniqueUser `json:"blog_findUniqueUser,omitempty"`
}

type N07oidc__fromclaim__fromclaimResponseData_blog_findUniqueUser struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N07oidc__fromclaim__fromclaim_createInput struct {
	Title string `json:"title,omitempty"`
}

type N07oidc__fromclaim__fromclaim_createInternalInput struct {
	Email string `json:"email"`
	Title string `json:"title,omitempty"`
}

type N07oidc__fromclaim__fromclaim_createResponseData struct {
	Blog_createOnePost N07oidc__fromclaim__fromclaim_createResponseData_blog_createOnePost `json:"blog_createOnePost,omitempty"`
}

type N07oidc__fromclaim__fromclaim_createResponseData_blog_createOnePost struct {
	User N07oidc__fromclaim__fromclaim_createResponseData_blog_createOnePost_User `json:"User,omitempty"`
	Id   int64                                                                    `json:"id"`
}

type N07oidc__fromclaim__fromclaim_createResponseData_blog_createOnePost_User struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

type N07oidc__fromclaim__fromclaim_updateInput struct {
	Id    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type N07oidc__fromclaim__fromclaim_updateInternalInput struct {
	Email string `json:"email"`
	Id    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
}

type N07oidc__fromclaim__fromclaim_updateResponseData struct {
	Blog_updateOnePost N07oidc__fromclaim__fromclaim_updateResponseData_blog_updateOnePost `json:"blog_updateOnePost,omitempty"`
}

type N07oidc__fromclaim__fromclaim_updateResponseData_blog_updateOnePost struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N10hook__fromclaim_customInput struct {
}

type N10hook__fromclaim_customInternalInput struct {
	Title string `json:"title"`
}

type N10hook__fromclaim_customResponseData struct {
	Todo_findManyTodo []N10hook__fromclaim_customResponseData_todo_findManyTodo `json:"todo_findManyTodo"`
}

type N10hook__fromclaim_customResponseData_todo_findManyTodo struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type N10hook__internalInput struct {
}

type N10hook__internalInternalInput struct {
}

type N10hook__internalResponseData struct {
	Blog_findFirstUser N10hook__internalResponseData_blog_findFirstUser `json:"blog_findFirstUser,omitempty"`
}

type N10hook__internalResponseData_blog_findFirstUser struct {
	Age     int64  `json:"age,omitempty"`
	Country string `json:"country,omitempty"`
	Email   string `json:"email"`
	Id      int64  `json:"id"`
}

type N10hook__rbacInput struct {
}

type N10hook__rbacInternalInput struct {
}

type N10hook__rbacResponseData struct {
	Blog_deleteManyUser N10hook__rbacResponseData_blog_deleteManyUser `json:"blog_deleteManyUser,omitempty"`
}

type N10hook__rbacResponseData_blog_deleteManyUser struct {
	Count int64 `json:"count"`
}

type N20directive__customizedfield__createOneInput struct {
	Content   string `json:"content"`
	Id        int64  `json:"id"`
	Published bool   `json:"published,omitempty"`
	Title     string `json:"title"`
}

type N20directive__customizedfield__createOneInternalInput struct {
	Content   string `json:"content"`
	Id        int64  `json:"id"`
	Published bool   `json:"published,omitempty"`
	Title     string `json:"title"`
}

type N20directive__customizedfield__createOneResponseData struct {
	Data N20directive__customizedfield__createOneResponseData_data `json:"data,omitempty"`
}

type N20directive__customizedfield__createOneResponseData_data struct {
	/* 内容是否合法 */
	ContentValid bool  `json:"contentValid,omitempty"`
	Id           int64 `json:"id"`
}

type N20directive__hookvarible__createOneInput struct {
	Content   string                  `json:"content"`
	Likes     []*Blog_UserCreateInput `json:"likes,omitempty"`
	Published bool                    `json:"published,omitempty"`
	Title     string                  `json:"title"`
}

type N20directive__hookvarible__createOneInternalInput struct {
	Content   string                  `json:"content"`
	Likes     []*Blog_UserCreateInput `json:"likes,omitempty"`
	Published bool                    `json:"published,omitempty"`
	Title     string                  `json:"title"`
}

type N20directive__hookvarible__createOneResponseData struct {
	Data N20directive__hookvarible__createOneResponseData_data `json:"data,omitempty"`
}

type N20directive__hookvarible__createOneResponseData_data struct {
	Id int64 `json:"id"`
}

type N20directive__queryraw__afterInput struct {
}

type N20directive__queryraw__afterInternalInput struct {
}

type N20directive__queryraw__afterResponseData struct {
	Blog_executeRaw any                                                       `json:"blog_executeRaw,omitempty"`
	Blog_queryRaw   []N20directive__queryraw__afterResponseData_blog_queryRaw `json:"blog_queryRaw,omitempty"`
}

type N20directive__queryraw__afterResponseData_blog_queryRaw struct {
	Age     int64  `json:"age,omitempty"`
	Country string `json:"country,omitempty"`
	Email   string `json:"email,omitempty"`
	Id      int64  `json:"id,omitempty"`
}

type N20directive__queryraw__beforeInput struct {
}

type N20directive__queryraw__beforeInternalInput struct {
}

type N20directive__queryraw__beforeResponseData struct {
	Blog_executeRaw any `json:"blog_executeRaw,omitempty"`
	Blog_queryRaw   any `json:"blog_queryRaw,omitempty"`
}

type N20directive__skip_include__createIfNotExistInput struct {
	CategoryName string `json:"categoryName"`
	Content      string `json:"content"`
	Title        string `json:"title"`
	UserId       int64  `json:"userId"`
}

type N20directive__skip_include__createIfNotExistInternalInput struct {
	CategoryName string `json:"categoryName"`
	Content      string `json:"content"`
	Title        string `json:"title"`
	UserId       int64  `json:"userId"`
}

type N20directive__skip_include__createIfNotExistResponseData struct {
	Data N20directive__skip_include__createIfNotExistResponseData_data `json:"data,omitempty"`
}

type N20directive__skip_include__createIfNotExistResponseData_data struct {
	Created N20directive__skip_include__createIfNotExistResponseData_data_created `json:"created"`
	Email   string                                                                `json:"email"`
	Existed N20directive__skip_include__createIfNotExistResponseData_data_existed `json:"existed"`
}

type N20directive__skip_include__createIfNotExistResponseData_data_created struct {
	CategoryId int64 `json:"categoryId,omitempty"`
	PostId     int64 `json:"postId,omitempty"`
}

type N20directive__skip_include__createIfNotExistResponseData_data_existed struct {
	CategoryId int64 `json:"categoryId,omitempty"`
}

type N20directive__transaction__recreateInput struct {
	Content   string `json:"content"`
	Published bool   `json:"published,omitempty"`
	Title     string `json:"title"`
}

type N20directive__transaction__recreateInternalInput struct {
	Content   string `json:"content"`
	Published bool   `json:"published,omitempty"`
	Title     string `json:"title"`
}

type N20directive__transaction__recreateResponseData struct {
	Data    N20directive__transaction__recreateResponseData_data `json:"data,omitempty"`
	Deleted int64                                                `json:"deleted,omitempty"`
}

type N20directive__transaction__recreateResponseData_data struct {
	Id int64 `json:"id"`
}
