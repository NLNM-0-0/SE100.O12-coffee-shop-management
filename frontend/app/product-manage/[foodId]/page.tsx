"use client";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Card, CardContent } from "@/components/ui/card";

import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useEffect, useState } from "react";
import {
  SubmitErrorHandler,
  SubmitHandler,
  useFieldArray,
  useForm,
} from "react-hook-form";
import { AiOutlineClose } from "react-icons/ai";
import { required } from "@/constants";
import CategoryList from "@/components/product-manage/category-list";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import SizeInsert from "@/components/stock-manage/size-ingredient-insert";
import { toast } from "@/components/ui/use-toast";
import ConfirmDialog from "@/components/confirm-dialog";
import { LuCheck } from "react-icons/lu";
import { FiTrash2 } from "react-icons/fi";
import { Textarea } from "@/components/ui/textarea";
import { imageUpload } from "@/lib/staff/uploadImage";
import { useRouter } from "next/navigation";
import Image from "next/image";
import getFood from "@/lib/food/getFood";
import Loading from "@/components/loading";
import SizeUpdate from "@/components/stock-manage/size-ingredient-update";
import updateFood from "@/lib/food/updateFood";
import { useCurrentUser } from "@/hooks/use-user";
import { includesRoles } from "@/lib/utils";
import NoRole from "@/components/no-role";
export const FormSchema = z.object({
  id: z.string().max(12, "Tối đa 12 ký tự"),
  name: required,
  description: z.string(),
  cookingGuide: z.string(),
  image: z.string(),
  categories: z
    .array(z.object({ idCate: z.string() }))
    .refine((category) => category.length > 0, {
      message: "Vui lòng chọn ít nhất một danh mục",
    }),
  sizes: z
    .array(
      z.object({
        sizeId: z.string(),
        name: required,
        cost: z.coerce
          .number()
          .int("Giá vốn phải là số nguyên")
          .gt(0, "Giá vốn phải lớn hơn 0"),
        price: z.coerce.number().gt(0, "Giá vốn phải lớn hơn 0"),
        recipe: z.object({
          details: z
            .array(
              z.object({
                ingredientId: z.string(),

                amountNeed: z.coerce.number(),
              })
            )
            .refine((value) => value.length > 0, {
              message: "Vui lòng có ít nhất một nguyên liệu",
            }),
        }),
      })
    )
    .refine((size) => size.length > 0, {
      message: "Vui lòng có ít nhất một kích cỡ",
    }),
});
const EditProductPage = ({ params }: { params: { foodId: string } }) => {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
  });
  const {
    register,
    handleSubmit,
    control,
    watch,
    reset,
    formState: { errors, isDirty },
  } = form;

  const { fields, append, remove, update } = useFieldArray({
    control: control,
    name: "sizes",
  });
  const {
    fields: fieldsCate,
    append: appendCate,
    remove: removeCate,
    update: updateCate,
  } = useFieldArray({
    control: control,
    name: "categories",
  });
  const sizes = watch("sizes");
  const router = useRouter();
  const { data, isLoading, isError, mutate } = getFood(params.foodId);

  const onFormSubmit: SubmitHandler<z.infer<typeof FormSchema>> = async (
    data
  ) => {
    console.log(data);
    if (image) {
      let formData = new FormData();

      formData.append("file", image);
      formData.append("imageType", "Food");

      const imgRes = await imageUpload(formData);
      if (imgRes.hasOwnProperty("errorKey")) {
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: imgRes.message,
        });
        return;
      }

      data.image = imgRes.data;
    }

    const food = {
      id: data.id,
      name: data.name,
      description: data.description,
      cookingGuide: data.cookingGuide,
      image: data.image,
      categories: data.categories.map((item) => item.idCate),
      sizes: data.sizes,
    };
    const response: Promise<any> = updateFood({ food: food });
    const responseData = await response;

    if (responseData.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      toast({
        variant: "success",
        title: "Thành công",
        description: "Chỉnh sửa sản phẩm thành công",
      });
      mutate();
      router.refresh();
    }
  };
  const onErrors: SubmitErrorHandler<z.infer<typeof FormSchema>> = (data) => {
    toast({
      variant: "destructive",
      title: "Có lỗi",
      description: "Vui lòng điền đầy đủ thông tin",
    });
  };
  const [image, setImage] = useState<any>();
  const [imagePreviews, setImagePreviews] = useState<any>();
  const handleMultipleImage = (event: any) => {
    const file = event.target.files[0];
    if (file) {
      if (file && file.type.includes("image")) {
        setImage(file);
        console.log(file.type);
        const reader = new FileReader();
        reader.onload = () => {
          setImagePreviews(reader.result);
        };
        reader.readAsDataURL(file);
      } else {
        setImage(null);
        toast({
          variant: "destructive",
          title: "Có lỗi",
          description: "File không hợp lệ",
        });
        console.log("file không hợp lệ");
      }
    } else {
      setImage(null);
    }
  };
  const handleReset = () => {
    reset({
      id: data.id,
      name: data.name,
      cookingGuide: data.cookingGuide,
      image: data.image,
      categories: data.categories.map((item) => {
        return { idCate: item.category.id };
      }),
      sizes: data.sizes.map((item) => {
        return {
          sizeId: item.sizeId,
          name: item.name,
          cost: item.cost,
          price: item.price,
          recipe: {
            details: item.recipe.details.map((detail) => {
              return {
                ingredientId: detail.ingredient.id,
                amountNeed: detail.amountNeed,
              };
            }),
          },
        };
      }),
    });
  };
  useEffect(() => {
    if (data) {
      handleReset();
    }
  }, [data]);
  const { currentUser } = useCurrentUser();
  const canEdit =
    currentUser &&
    includesRoles({
      currentUser: currentUser,
      allowedFeatures: ["FOD_UP_INFO"],
    });
  if (isError) return <div>Failed to load</div>;
  else if (!currentUser || isLoading) {
    return <Loading />;
  } else if (
    currentUser &&
    !includesRoles({
      currentUser: currentUser,
      allowedFeatures: ["FOD_VIEW"],
    })
  ) {
    return <NoRole></NoRole>;
  } else
    return (
      <div className="col items-center">
        <div className="col w-full xl:px-16">
          <div className="flex justify-between">
            <h1 className="font-medium text-xxl self-start">
              Chỉnh sửa sản phẩm
            </h1>
            {canEdit ? (
              <div className="flex md:justify-end justify-stretch gap-2">
                <Button
                  className="px-4 bg-white"
                  disabled={!isDirty}
                  variant={"outline"}
                  type="button"
                  onClick={() => {
                    if (data) {
                      handleReset();
                    }
                  }}
                >
                  <div className="flex flex-wrap gap-2 items-center">
                    <FiTrash2 className="text-muted-foreground" />
                    Hủy
                  </div>
                </Button>
                <ConfirmDialog
                  title="Xác nhận"
                  description="Bạn xác nhận chỉnh sửa mặt hàng ?"
                  handleYes={() => handleSubmit(onFormSubmit, onErrors)()}
                >
                  <Button className="px-4 pl-2">
                    <div className="flex flex-wrap gap-2 items-center">
                      <LuCheck />
                      Chỉnh sửa
                    </div>
                  </Button>
                </ConfirmDialog>
              </div>
            ) : null}
          </div>
          <form>
            <div className="flex flex-col gap-4 xl:flex-row">
              <div className="xl:basis-3/5 flex flex-col gap-4">
                <Card>
                  <CardContent className="p-6">
                    <div className="flex flex-col gap-4 2xl:flex-row 2xl:gap-2">
                      <div className="flex-1">
                        <Label htmlFor="prodName">Tên sản phẩm</Label>
                        <Input
                          id="prodName"
                          {...register(`name` as const)}
                        ></Input>
                        {errors.name && (
                          <span className="error___message">
                            {errors.name.message}
                          </span>
                        )}
                      </div>
                    </div>
                  </CardContent>
                </Card>
                <Card>
                  <CardContent className="p-6 pt-3">
                    <div className="flex flex-col gap-3">
                      {/* price list */}
                      {fields.map((field, index) => {
                        return (
                          <div key={field.id} className="flex gap-3">
                            <div className={`flex-1`}>
                              <Label>Tên giá</Label>
                              <Input
                                type="text"
                                {...register(`sizes.${index}.name` as const)}
                              ></Input>
                              {errors &&
                              errors.sizes &&
                              errors.sizes[index] &&
                              (errors.sizes[index]!.name as
                                | { message: string }
                                | undefined) ? (
                                <span className="error___message">
                                  {errors.sizes[index]!.name!.message}
                                </span>
                              ) : null}
                            </div>
                            <div className={`flex-1  `}>
                              <Label>Giá bán (VND)</Label>
                              <Input
                                {...register(`sizes.${index}.price` as const)}
                              ></Input>
                              {errors &&
                              errors.sizes &&
                              errors.sizes[index] &&
                              (errors.sizes[index]!.price as
                                | { message: string }
                                | undefined) ? (
                                <span className="error___message">
                                  {errors.sizes[index]!.price!.message}
                                </span>
                              ) : null}
                            </div>
                            <div className="flex-1">
                              <Label>Giá vốn (VND)</Label>
                              <Input
                                {...register(`sizes.${index}.cost` as const)}
                              ></Input>
                              {errors &&
                              errors.sizes &&
                              errors.sizes[index] &&
                              (errors.sizes[index]!.cost as
                                | { message: string }
                                | undefined) ? (
                                <span className="error___message">
                                  {errors.sizes[index]!.cost!.message}
                                </span>
                              ) : null}
                            </div>

                            {fields.length > 1 ? (
                              <Button
                                variant={"ghost"}
                                className={`self-end px-3 gap-0 `}
                                onClick={() => {
                                  remove(index);
                                }}
                              >
                                <AiOutlineClose />
                              </Button>
                            ) : (
                              <Button
                                variant={"ghost"}
                                className={`self-end px-3 gap-0 `}
                                disabled
                                onClick={() => {
                                  remove(index);
                                }}
                              >
                                <AiOutlineClose />
                              </Button>
                            )}
                          </div>
                        );
                      })}
                      <div>
                        <Button
                          className="self-start p-2"
                          variant={"link"}
                          type="button"
                          onClick={() => {
                            append({
                              sizeId: "",
                              price: 0,
                              cost: 0,
                              name: "",
                              recipe: { details: [] },
                            });
                          }}
                        >
                          <span className="font-bold">+</span> Thêm giá
                        </Button>
                      </div>
                    </div>
                  </CardContent>
                </Card>
                {/* ingredient list */}
                <Card>
                  <CardContent className="p-6 flex flex-col gap-4">
                    {/* ingredients list  */}

                    {
                      <Tabs defaultValue={"0"}>
                        <TabsList className="w-full justify-start mb-2 h-fit flex-wrap bg-white border-b pb-0">
                          {sizes.map((size, index) => (
                            <TabsTrigger
                              key={index}
                              className="tab___trigger"
                              value={index.toString()}
                            >
                              {size.name || "Tên giá "}
                            </TabsTrigger>
                          ))}
                        </TabsList>

                        {sizes.map((size, index) => (
                          <TabsContent key={index} value={index.toString()}>
                            <SizeUpdate form={form} sizeIndex={index} />
                          </TabsContent>
                        ))}
                      </Tabs>
                    }
                  </CardContent>
                </Card>
              </div>

              <Card className="xl:basis-2/5 xl:self-start">
                <CardContent className="p-6">
                  <div className="flex gap-4 flex-col ">
                    {/* category list */}
                    <div>
                      <Label>Danh mục</Label>
                      <CategoryList
                        isEdit
                        canAdd
                        checkedCategory={fieldsCate.map((cate) => cate.idCate)}
                        onCheckChanged={(idCate) => {
                          const selectedIndex = fieldsCate.findIndex(
                            (cate) => cate.idCate === idCate
                          );
                          if (selectedIndex > -1) {
                            removeCate(selectedIndex);
                          } else {
                            appendCate({ idCate: idCate });
                          }
                        }}
                        onRemove={(index) => {
                          removeCate(index);
                        }}
                      />
                      {errors.categories && (
                        <span className="error___message">
                          {errors.categories.message}
                        </span>
                      )}
                    </div>
                    <div>
                      <Label>Mô tả</Label>
                      <Textarea {...register("description")} />
                    </div>
                    <div>
                      <Label>Công thức nấu</Label>
                      <Textarea {...register("cookingGuide")} />
                    </div>
                    <div>
                      <Label htmlFor="img">Hình ảnh</Label>
                      <div className="flex items-center gap-4">
                        <Input
                          className="basis-1/2"
                          id="img"
                          type="file"
                          onChange={handleMultipleImage}
                        ></Input>
                        <div>
                          {image && imagePreviews ? (
                            <img
                              src={imagePreviews}
                              alt={`Preview`}
                              className="h-[120px] w-auto object-cover"
                            />
                          ) : (
                            <Image
                              src={data.image ?? "/no-image.jpg"}
                              alt="ảnh"
                              className="h-[120px] w-auto  object-cover"
                              height={120}
                              width={120}
                            ></Image>
                          )}
                        </div>
                      </div>
                    </div>
                  </div>
                </CardContent>
              </Card>
            </div>
          </form>
        </div>
      </div>
    );
};

export default EditProductPage;
